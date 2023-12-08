package api

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	var request request.GetPhotoRequest
	request.PathParameters.PhotoId = photoId

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve author of the photo
	ctx.Logger.Debugf("retrieving photo authorId from database")
	authorId, err := rt.db.GetPhotoAuthorId(photoId)
	if err == sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("error retrieving photo authorId from database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photo authorId from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if logged in user is banned by author of the photo
	ctx.Logger.Debugf(`checking if ban (bannerId: "%s", bannedId "%s") exists in database`, authorId, ctx.User.UserId)
	if banned, err := rt.db.ExistsBan(authorId, ctx.User.UserId); err != nil {
		ctx.Logger.WithError(err).Error("error searching ban in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banned {
		ctx.Logger.Error("requested user is banned by logged in user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Retrieve photo
	ctx.Logger.Debugf("retrieving photo from file system")
	photoPath, err := rt.fs.GetPhotoPath(photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photo from file system")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Serve photo
	w.Header().Set("content-type", "img/png")
	http.ServeFile(w, r, photoPath)
}