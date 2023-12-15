package request

type UncommentPhotoRequest struct {
	PathParameters struct {
		PhotoID   string `key:"photoID"`
		CommentID string `key:"commentID"`
	}
}

func (request *UncommentPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) &&
		isValidUUID(request.PathParameters.CommentID)
}
