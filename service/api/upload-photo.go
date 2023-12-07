package api

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

const (
	maxSize          = 10 << 20
	imageFormat      = "image/png"
	maxCaptionLenght = 2200
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	r.ParseMultipartForm(maxSize)
	// Retrieving file
	ctx.Logger.Debugf("retrieving file and caption")
	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	caption := r.FormValue("caption")

	// Validating file
	if fileHeader.Header.Get("Content-Type") != imageFormat {
		ctx.Logger.WithError(err).Error("error validating file")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	if fileHeader.Size > maxSize {
		ctx.Logger.WithError(err).Error("error validating file")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	// Validating caption
	if len(caption) > 2200 {
		ctx.Logger.WithError(err).Error("error validating caption")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Creating new photo UUID
	photoUUID, err := uuid.NewV4()
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating new UUID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photoId := photoUUID.String()

	// Saving photo to filesystem
	ctx.Logger.Debugf("saving photo to filesystem")
	err = rt.fs.SavePhoto(&file, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving photo to filesystem")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Saving record to database
	ctx.Logger.Debugf("saving photo record to database")
	dateTime := globaltime.ToString(globaltime.Now())
	err = rt.db.InsertPhoto(photoId, ctx.User.UserId, caption, dateTime)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving photo record to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := response.UploadPhotoResponse{PhotoId: photoId}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
