package request

type CommentPhotoRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
	Text string `json:"text"`
}

func (request *CommentPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) && isValidText(request.Text)
}
