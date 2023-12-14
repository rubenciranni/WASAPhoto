package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getFollowing(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userId := ps.ByName("userId")
	startId := ps.ByName("startId")
	var req request.GetFollowingRequest
	req.PathParameters.UserId = userId
	req.QueryParameters.StartId = startId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if logged-in user is banned by requested user
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, req.PathParameters.UserId, ctx.User.UserId)
	if banned, err := rt.db.ExistsBan(userId, ctx.User.UserId); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve following from database
	ctx.Logger.Debugf(`retrieving following of "%s" from database`, userId)
	following, err := rt.db.GetFollowing(userId, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving following from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.GetFollowingResponse
	if len(following) == 0 {
		res.LastId = ""
	} else {
		res.LastId = following[len(following)-1].UserId
	}
	res.Records = following
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
