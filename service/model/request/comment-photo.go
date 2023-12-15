package request

type CommentPhotoRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
	Text string `json:"text"`
}

func (request *CommentPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId) && isValidText(request.Text)
}
