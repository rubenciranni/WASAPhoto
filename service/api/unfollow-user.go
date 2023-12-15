package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userID := ps.ByName("userID")
	var req request.UnfollowUserRequest
	req.PathParameters.UserID = userID

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if logged-in user is the requested user
	if ctx.User.UserID == userID {
		ctx.Logger.Error("error: user is trying to unfollow himself")
		w.WriteHeader(http.StatusForbidden)
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

	// Delete follow from database
	ctx.Logger.Debugf(`Deleting follow (followerId: "%s", followedId: "%s") into database`, ctx.User.UserID, userID)
	err := rt.db.DeleteFollow(ctx.User.UserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error inserting follow into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
