package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	var req request.SetMyUserNameRequest
	ctx.Logger.Debugf("decoding JSON")
	err := json.NewDecoder(r.Body).Decode(&req)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if newUsername is already taken by another user
	ctx.Logger.Debugf(`checking username "%s" availability`, req.NewUsername)
	if userID, err := rt.db.GetUserID(req.NewUsername); err == nil && userID != ctx.User.UserID {
		ctx.Logger.Errorf(`username "%s" is already taken`, req.NewUsername)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Update username
	ctx.Logger.Debugf("updating username")
	err = rt.db.SetUserName(ctx.User.UserID, req.NewUsername)
	if err != nil {
		ctx.Logger.WithError(err).Error("error updating username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
