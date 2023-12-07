package request

type CommentPhotoRequest struct {
	Text string `json:"text"`
}

const maxTextLenght = 2200

func (request *CommentPhotoRequest) IsValid() bool {
	return len(request.Text) < maxTextLenght
}
