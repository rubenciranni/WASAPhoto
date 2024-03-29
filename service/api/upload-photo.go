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
	maxCaptionLength = 2200
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx.Logger.Debugf("retrieving file and caption")
	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	caption := r.FormValue("caption")

	// Validate file
	if fileHeader.Header.Get("Content-Type") != imageFormat {
		ctx.Logger.Error("error validating file")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	if fileHeader.Size > maxSize {
		ctx.Logger.Error("error validating file")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	// Validate caption
	if len(caption) > maxCaptionLength {
		ctx.Logger.Error("error validating caption")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new photo UUID
	photoUUID, err := uuid.NewV4()
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating new UUID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photoId := photoUUID.String()

	// Save photo to file system
	ctx.Logger.Debugf("saving photo to file system")
	err = rt.fs.SavePhoto(&file, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving photo to file system")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save record to database
	ctx.Logger.Debugf("inserting photo record into database")
	dateTime := globaltime.ToString(globaltime.Now())
	err = rt.db.InsertPhoto(photoId, ctx.User.UserId, caption, dateTime)
	if err != nil {
		ctx.Logger.WithError(err).Error("error inserting photo record into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	res := response.UploadPhotoResponse{PhotoId: photoId}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}
