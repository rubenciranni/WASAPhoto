package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	var request request.SetMyUserNameRequest
	ctx.Logger.Debugf("deconding JSON")
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("error deconding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if newUsername is already taken by another user
	ctx.Logger.Debugf(`checking username "%s" availability`, request.NewUsername)
	if userId, err := rt.db.GetUserId(request.NewUsername); err == nil && userId != ctx.User.UserId {
		ctx.Logger.Errorf(`username "%s" is already taken`, request.NewUsername)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Update username
	ctx.Logger.Debugf("updating username")
	err = rt.db.SetUserName(ctx.User.UserId, request.NewUsername)
	if err != nil {
		ctx.Logger.WithError(err).Error("error updating username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
