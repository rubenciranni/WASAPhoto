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

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	var req request.DoLoginRequest
	ctx.Logger.Debugf("decoding JSON")
	err := json.NewDecoder(r.Body).Decode(&req)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check username existance
	ctx.Logger.Debugf(`retrieving userID for user "%s"`, req.Username)
	userID, err := rt.db.GetUserID(req.Username)

	// Create new user if username doesn't exist
	if errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.Debugf(`user "%s" does not exist, creating new user`, req.Username)
		newUserID, err := uuid.NewV4()
		if err != nil {
			ctx.Logger.WithError(err).Error("error creating new UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userID = newUserID.String()
		err = rt.db.InsertUser(userID, req.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("error inserting new user in database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error getting userID from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	res := response.DoLoginResponse{UserID: userID}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
