package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	var request request.DoLoginRequest
	ctx.Logger.Debugf("deconding JSON")
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check username existance
	ctx.Logger.Debugf(`retrieving userId for user "%s"`, request.Username)
	userId, err := rt.db.GetUserId(request.Username)

	// Create new user if username doesn't exists
	if errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.Debugf(`user "%s" does not exist, creating new user`, request.Username)
		newUserId, err := uuid.NewV4()
		if err != nil {
			ctx.Logger.WithError(err).Error("error creating new UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userId = newUserId.String()
		err = rt.db.InsertUser(userId, request.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("error inserting new user in database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error getting userId from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	response := response.DoLoginResponse{UserId: userId}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
