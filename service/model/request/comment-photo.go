package request

type CommentPhotoRequest struct {
	Text string `json:"text"`
}

func (request *CommentPhotoRequest) IsValid() bool {
	return 0 < len(request.Text) && len(request.Text) < 2200
}
