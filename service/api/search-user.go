package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	username := r.URL.Query().Get("username")
	startID := r.URL.Query().Get("startID")
	var req request.SearchUserRequest
	req.QueryParameters.Username = username
	req.QueryParameters.StartID = startID

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve users from database
	ctx.Logger.Debug("retrieving users from database")
	users, err := rt.db.GetUsers(ctx.User.UserID, username, startID)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("error retrieving users from database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving users from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.SearchUserResponse
	if len(users) == 0 {
		res.LastID = ""
	} else {
		res.LastID = users[len(users)-1].UserID
	}
	res.Records = users
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
