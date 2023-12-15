package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse req
	var req request.CommentPhotoRequest
	ctx.Logger.Debugf("decoding JSON")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = r.Body.Close()
	photoID := ps.ByName("photoID")
	req.PathParameters.PhotoID = photoID

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve author of the photo
	ctx.Logger.Debugf("retrieving photo authorId from database")
	authorId, err := rt.db.GetPhotoAuthorId(photoID)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("error retrieving photo authorId from database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photo authorId from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if logged-in user is banned by author of the photo
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, authorId, ctx.User.UserID)
	if banned, err := rt.db.ExistsBan(authorId, ctx.User.UserID); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Generate a new comment UUID
	commentUUID, err := uuid.NewV4()
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating new UUID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	commentId := commentUUID.String()

	// Insert comment into database
	ctx.Logger.Debugf("inserting comment into database")
	dateTime := globaltime.ToString(globaltime.Now())
	err = rt.db.InsertComment(commentId, photoID, ctx.User.UserID, req.Text, dateTime)
	if err != nil {
		ctx.Logger.WithError(err).Error("error inserting comment into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	res := response.CommentPhotoResponse{CommentID: commentId}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
