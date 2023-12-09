package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	username := r.URL.Query().Get("username")
	startId := r.URL.Query().Get("startId")
	var request request.SearchUserRequest
	request.QueryParameters.Username = username
	request.QueryParameters.StartId = startId

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve users from database
	ctx.Logger.Debug("retrieving users from database")
	users, err := rt.db.GetUsers(ctx.User.UserId, username, startId)
	if err == sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("error retrieving users from database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving users from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var response response.SearchUserResponse
	if len(users) == 0 {
		response.LastId = ""
	} else {
		response.LastId = users[len(users)-1].UserId
	}
	response.Records = users
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
