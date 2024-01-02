package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	var req request.GetPhotoRequest
	req.PathParameters.PhotoId = photoId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
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
