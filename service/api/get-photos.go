package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userId := r.URL.Query().Get("userId")
	startDate := r.URL.Query().Get("startDate")
	startId := r.URL.Query().Get("startId")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var request request.GetPhotosRequest
	request.QueryParameters.UserId = userId
	request.QueryParameters.StartDate = startDate
	request.QueryParameters.StartId = startId

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if provided userId exists
	ctx.Logger.Debugf(`retrieving user for userId "%s"`, request.QueryParameters.UserId)
	if _, err := rt.db.GetUser(userId); err == sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("userId does not exist in database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error searching userId in database")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Check if logged in user is banned by requested user
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, request.QueryParameters.UserId, ctx.User.UserId)
	if banned, err := rt.db.ExistsBan(userId, ctx.User.UserId); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve photos from database
	ctx.Logger.Debugf(`retrieving photos by userId "%s"`, request.QueryParameters.UserId)
	photos, err := rt.db.GetPhotosByUser(userId, startDate, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photos from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var response response.GetPhotosResponse
	if len(photos) == 0 {
		response.LastDate = ""
		response.LastId = ""
	} else {
		response.LastDate = photos[len(photos)-1].DateTime
		response.LastId = photos[len(photos)-1].PhotoId
	}
	response.Records = photos
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
