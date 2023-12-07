package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parsing request
	var request request.DoLoginRequest
	ctx.Logger.Debugf("deconding JSON")
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	regexpPattern := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)
	if !regexpPattern.MatchString(request.Username) {
		ctx.Logger.Error("error validating JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Checking username existance
	ctx.Logger.Debugf(`retrieving userId for user "%s"`, request.Username)
	userId, err := rt.db.GetUserId(request.Username)

	// Creating new user if username doesn't exists
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

	// Creating response
	response := response.DoLoginResponse{UserId: userId}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		ctx.Logger.WithError(err).Error("error encoding response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
}
