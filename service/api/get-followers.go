package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getFollowers(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userID := ps.ByName("userID")
	startId := ps.ByName("startId")
	var req request.GetFollowersRequest
	req.PathParameters.UserID = userID
	req.QueryParameters.StartId = startId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if logged-in user is banned by requested user
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, req.PathParameters.UserID, ctx.User.UserID)
	if banned, err := rt.db.ExistsBan(userID, ctx.User.UserID); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve followers from database
	ctx.Logger.Debugf(`retrieving followers of "%s" from database`, userID)
	followers, err := rt.db.GetFollowers(userID, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving followers from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.GetFollowersResponse
	if len(followers) == 0 {
		res.LastId = ""
	} else {
		res.LastId = followers[len(followers)-1].UserID
	}
	res.Records = followers
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
