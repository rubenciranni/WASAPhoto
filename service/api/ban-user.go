package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) banUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userId := ps.ByName("userId")
	var req request.BanUserRequest
	req.PathParameters.UserId = userId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if logged-in user is the requested user
	if ctx.User.UserId == userId {
		ctx.Logger.Error("error: user is trying to ban himself")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Insert ban into database
	ctx.Logger.Debugf(`Inserting ban (bannerId: "%s", bannedId: "%s") into database`, ctx.User.UserId, userId)
	err := rt.db.InsertBan(ctx.User.UserId, userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error inserting ban into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
