package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	startId := r.URL.Query().Get("startId")
	var request request.GetLikesRequest
	request.PathParameters.PhotoId = photoId
	request.QueryParameters.StartId = startId

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

	// Get likes from database
	ctx.Logger.Debugf(`retrieving likes for photoId "%s"`, request.PathParameters.PhotoId)
	likes, err := rt.db.GetLikes(photoId, startId)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving likes from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	response := response.GetLikesResponse{}
	if len(likes) == 0 {
		response.LastId = ""
	} else {
		response.LastId = likes[len(likes)-1].UserId
	}
	response.Records = likes
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
