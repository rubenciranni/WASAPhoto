package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rubenciranni/WASAPhoto/service/api/reqcontext"
	"github.com/rubenciranni/WASAPhoto/service/globaltime"
	"github.com/rubenciranni/WASAPhoto/service/model/request"
	"github.com/rubenciranni/WASAPhoto/service/model/response"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoID := ps.ByName("photoID")
	startID := r.URL.Query().Get("startID")
	startDate := r.URL.Query().Get("startDate")
	if startDate == "" {
		startDate = globaltime.ToString(globaltime.Now())
	}
	var req request.GetCommentsRequest
	req.PathParameters.PhotoID = photoID
	req.QueryParameters.StartID = startID
	req.QueryParameters.StartDate = startDate

	// Validate request
	ctx.Logger.Debug(photoID)
	ctx.Logger.Debug(startID)
	ctx.Logger.Debug(startDate)
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

	// Get comments from database
	ctx.Logger.Debugf(`retrieving comments for photoID "%s"`, req.PathParameters.PhotoID)
	comments, err := rt.db.GetComments(photoID, startDate, startID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving comments from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	var res response.GetCommentsResponse
	if len(comments) == 0 {
		res.LastDate = ""
		res.LastID = ""
	} else {
		res.LastDate = comments[len(comments)-1].DateTime
		res.LastID = comments[len(comments)-1].CommentID
	}
	res.Records = comments
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}
