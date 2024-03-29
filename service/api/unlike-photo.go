package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	var req request.LikePhotoRequest
	req.PathParameters.PhotoId = photoId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve author of the photo
	ctx.Logger.Debugf("retrieving photo authorId from database")
	authorId, err := rt.db.GetPhotoAuthorId(photoId)
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
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, authorId, ctx.User.UserId)
	if banned, err := rt.db.ExistsBan(authorId, ctx.User.UserId); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged-in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete Like
	ctx.Logger.Debugf(`Deleting like from photo "%s"`, photoId)
	err = rt.db.DeleteLike(photoId, ctx.User.UserId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error deleting like into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
