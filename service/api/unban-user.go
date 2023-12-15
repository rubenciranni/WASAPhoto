package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) unbanUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userID := ps.ByName("userID")
	var req request.UnbanUserRequest
	req.PathParameters.UserID = userID

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if logged-in user is the requested user
	if ctx.User.UserID == userID {
		ctx.Logger.Error("error: user is trying to unban himself")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Deleting ban from database
	ctx.Logger.Debugf(`Deleting ban (bannerId: "%s", bannedId: "%s") from database`, ctx.User.UserID, userID)
	err := rt.db.DeleteBan(ctx.User.UserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error deleting ban into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
