package request

type UncommentPhotoRequest struct {
	PathParameters struct {
		PhotoId   string `key:"photoId"`
		CommentId string `key:"commentId"`
	}
}

func (request *UncommentPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId) &&
		isValidUUID(request.PathParameters.CommentId)
}
