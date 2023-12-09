package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	startDate := r.URL.Query().Get("startDate")
	startId := r.URL.Query().Get("startId")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var request request.GetMyStreamRequest
	request.QueryParameters.StartDate = startDate
	request.QueryParameters.StartId = startId

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve photos from database
	ctx.Logger.Debugf(`retrieving stream of "%s"`, ctx.User.UserId)
	photos, err := rt.db.GetStream(ctx.User.UserId, startDate, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photos from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var response response.GetMyStreamResponse
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
