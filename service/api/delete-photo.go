package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	var req request.DeletePhotoRequest
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

	// Check if logged-in user is the author of the photo
	ctx.Logger.Debugf("checking if logged-in user is the author of the photo")
	if authorId != ctx.User.UserId {
		ctx.Logger.Error("error: logged-in user is not the author of the photo")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete photo from file system
	ctx.Logger.Debugf("deleting photo from filesystem")
	err = rt.fs.DeletePhoto(photoId)
	if err != nil {
		ctx.Logger.Error("error deleting photo from file system")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Delete photo from database
	ctx.Logger.Debugf("deleting photo from database")
	err = rt.db.DeletePhoto(photoId)
	if err != nil {
		ctx.Logger.Error("error deleting photo from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	res := response.DeletePhotoResponse{PhotoId: photoId}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
