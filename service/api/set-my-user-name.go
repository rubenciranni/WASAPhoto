package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parsing request
	var request request.SetMyUserNameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse request body for setMyUserName operation")
		w.Header().Set("content-type", "application/json")
		response := response.Problem{
			Title:  "Bad Request",
			Status: http.StatusBadRequest,
			Detail: "Cannot parse request body."}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			ctx.Logger.WithError(err).Error("can't encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Debug("sending response")
		return
	}
	ctx.Logger.Debugf(`received and parsed request for setMyUserName operation`)

	// Checking if newUsername is already taken by another user
	if userId, err := rt.db.GetUserId(request.NewUsername); err == nil && userId != ctx.User.UserId {
		ctx.Logger.Debugf(`username "%s" is already taken, responding forbidden`, request.NewUsername)
		w.Header().Set("content-type", "application/json")
		response := response.Problem{
			Title:  "Forbidden",
			Status: http.StatusForbidden,
			Detail: "Username is already taken."}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			ctx.Logger.WithError(err).Error("can't encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Debug("sending response")
		return
	}

	// Updating username
	err = rt.db.SetUserName(ctx.User.UserId, request.NewUsername)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't update username")
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := response.Problem{
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: "Cannot update username"}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			ctx.Logger.WithError(err).Error("can't encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Logger.Debug("sending response")
		return
	}

	// Creating response
	w.WriteHeader(http.StatusNoContent)
	ctx.Logger.Debug("sending response")
}
