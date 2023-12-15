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

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	startDate := r.URL.Query().Get("startDate")
	startId := r.URL.Query().Get("startId")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var req request.GetMyStreamRequest
	req.QueryParameters.StartDate = startDate
	req.QueryParameters.StartId = startId

	// Validate request
	if !req.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve photos from database
	ctx.Logger.Debugf(`retrieving stream of "%s"`, ctx.User.UserID)
	photos, err := rt.db.GetStream(ctx.User.UserID, startDate, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving photos from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send res
	var res response.GetMyStreamResponse
	if len(photos) == 0 {
		res.LastDate = ""
		res.LastId = ""
	} else {
		res.LastDate = photos[len(photos)-1].DateTime
		res.LastId = photos[len(photos)-1].PhotoID
	}
	res.Records = photos
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
