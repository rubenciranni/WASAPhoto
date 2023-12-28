package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
	"github.com/rubenciranni/WASAPhoto/service/model/schema"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	username := r.URL.Query().Get("username")
	rawValue := r.URL.Query().Get("isExactMatch")
	var isExactMatch bool
	if rawValue == "" {
		isExactMatch = false
	} else {
		var err error
		isExactMatch, err = strconv.ParseBool(rawValue)
		if err != nil {
			ctx.Logger.WithError(err).Error("error converting string to boolean")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	startId := r.URL.Query().Get("startId")
	var req request.SearchUserRequest
	req.QueryParameters.Username = username
	req.QueryParameters.StartId = startId
	req.QueryParameters.IsExactMatch = isExactMatch

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve users from database
	ctx.Logger.Debug("retrieving users from database")
	var users []schema.User

	if isExactMatch {
		// Check if user exists
		userId, err := rt.db.GetUserId(username)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.Logger.WithError(err).Error("error retrieving users from database")
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			ctx.Logger.WithError(err).Error("error getting userId from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Check if logged-in user is banned by user
		ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, userId, ctx.User.UserId)
		if banned, err := rt.db.ExistsBan(userId, ctx.User.UserId); err != nil {
			ctx.Logger.WithError(err).Error("error searching ban in database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if banned {
			ctx.Logger.Error("requested user is banned by logged-in user")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user := schema.User{Username: username, UserId: userId}
		users = append(users, user)

	} else {
		var err error
		users, err = rt.db.GetUsers(ctx.User.UserId, username, startId)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.Logger.WithError(err).Error("error retrieving users from database")
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			ctx.Logger.WithError(err).Error("error retrieving users from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Send response
	var res response.SearchUserResponse
	if len(users) == 0 {
		res.LastId = ""
	} else {
		res.LastId = users[len(users)-1].UserId
	}
	res.Records = users
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
