package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	userId := r.URL.Query().Get("userId")
	startDate := r.URL.Query().Get("startDate")
	startId := r.URL.Query().Get("startId")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var req request.GetPhotosRequest
	req.QueryParameters.UserId = userId
	req.QueryParameters.StartDate = startDate
	req.QueryParameters.StartId = startId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if provided userId exists
	ctx.Logger.Debugf(`retrieving user for userId "%s"`, req.QueryParameters.UserId)
	if _, err := rt.db.GetUser(userId); errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("userId does not exist in database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error searching userId in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if logged-in user is banned by requested user
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, req.QueryParameters.UserId, ctx.User.UserId)
	if banned, err := rt.db.ExistsBan(userId, ctx.User.UserId); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve photos from database
	ctx.Logger.Debugf(`retrieving photos by userId "%s"`, req.QueryParameters.UserId)
	photos, err := rt.db.GetPhotosByUser(userId, startDate, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photos from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.GetPhotosResponse
	if len(photos) == 0 {
		res.LastDate = ""
		res.LastId = ""
	} else {
		res.LastDate = photos[len(photos)-1].DateTime
		res.LastId = photos[len(photos)-1].PhotoId
	}
	res.Records = photos
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
