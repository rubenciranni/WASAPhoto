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
	userID := r.URL.Query().Get("userID")
	startDate := r.URL.Query().Get("startDate")
	startID := r.URL.Query().Get("startID")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var req request.GetPhotosRequest
	req.QueryParameters.UserID = userID
	req.QueryParameters.StartDate = startDate
	req.QueryParameters.StartID = startID

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if provided userID exists
	ctx.Logger.Debugf(`retrieving user for userID "%s"`, req.QueryParameters.UserID)
	if _, err := rt.db.GetUser(userID); errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("userID does not exist in database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error searching userID in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if logged-in user is banned by requested user
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, req.QueryParameters.UserID, ctx.User.UserID)
	if banned, err := rt.db.ExistsBan(userID, ctx.User.UserID); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve photos from database
	ctx.Logger.Debugf(`retrieving photos by userID "%s"`, req.QueryParameters.UserID)
	photos, err := rt.db.GetPhotosByUser(userID, startDate, startID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photos from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.GetPhotosResponse
	if len(photos) == 0 {
		res.LastDate = ""
		res.LastID = ""
	} else {
		res.LastDate = photos[len(photos)-1].DateTime
		res.LastID = photos[len(photos)-1].PhotoID
	}
	res.Records = photos
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
