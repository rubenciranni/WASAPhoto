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

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse request
	photoId := ps.ByName("photoId")
	commentId := ps.ByName("commentId")
	var request request.UncommentPhotoRequest
	request.PathParameters.PhotoId = photoId
	request.PathParameters.CommentId = commentId

	// Validate request
	if !request.IsValid() {
		ctx.Logger.Error("error validating request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve author of the comment
	ctx.Logger.Debugf("retrieving comment authorId from database")
	authorId, err := rt.db.GetCommentAuthorId(commentId)
	if err == sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("error retrieving comment authorId from database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("error retrieving comment authorId from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if logged in user is the author of the comment
	ctx.Logger.Debugf("checking if logged in user is the author of the comment")
	if authorId != ctx.User.UserId {
		ctx.Logger.Error("error: logged in user is not the author of the comment")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete comment from database
	ctx.Logger.Debugf("deleting comment from database")
	err = rt.db.DeleteComment(commentId)
	if err != nil {
		ctx.Logger.Error("error deleting comment from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
	response := response.UncommentPhotoResponse{CommentId: commentId}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
