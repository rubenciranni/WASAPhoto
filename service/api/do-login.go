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
	w.Header().Set("content-type", "application/json")
	// Parsing request
	var request request.DoLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't parse request body for DoLogin operation")
		w.WriteHeader(http.StatusBadRequest)
		response := response.Problem{
			Title:  "Bad Request",
			Status: 400,
			Detail: "Cannot parse request body."}
		json.NewEncoder(w).Encode(response)
		ctx.Logger.Debug("sending response")
		return
	}
	ctx.Logger.Debugf(`received and parsed request for doLogin operation`)

	// Checking username validity
	regexpPattern := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)
	if !regexpPattern.MatchString(request.Username) {
		ctx.Logger.Debugf(`username in request is invalid`)
		w.WriteHeader(http.StatusBadRequest)
		response := response.Problem{
			Title:  "Bad Request",
			Status: 400,
			Detail: "Invalid username."}
		json.NewEncoder(w).Encode(response)
		ctx.Logger.Debug("sending response")
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
			ctx.Logger.WithError(err).Error("can't generate a new user UUID")
			w.WriteHeader(http.StatusInternalServerError)
			response := response.Problem{
				Title:  "Internal Server Error",
				Status: 500,
				Detail: "Cannot generate a new user UUID"}
			json.NewEncoder(w).Encode(response)
			ctx.Logger.Debug("sending response")
			return
		}
		userId = newUserId.String()
		err = rt.db.InsertUser(userId, request.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't insert a new user in database")
			w.WriteHeader(http.StatusInternalServerError)
			response := response.Problem{
				Title:  "Internal Server Error",
				Status: 500,
				Detail: "Cannot insert a new user in database"}
			json.NewEncoder(w).Encode(response)
			ctx.Logger.Debug("sending response")
			return
		}
	} else if err != nil {
		ctx.Logger.WithError(err).Error("can't get the userId for username")
		w.WriteHeader(http.StatusInternalServerError)
		response := response.Problem{
			Title:  "Internal Server Error",
			Status: 500,
			Detail: "Cannot get userId for requested username"}
		json.NewEncoder(w).Encode(response)
		ctx.Logger.Debug("sending response")
		return
	}

	// Creating response
	w.WriteHeader(http.StatusCreated)
	response := response.DoLoginResponse{UserId: userId}

	json.NewEncoder(w).Encode(response)
	ctx.Logger.Debug("sending response")
}
