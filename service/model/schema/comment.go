package schema

type Comment struct {
	CommentId string `json:"commentId"`
	Text      string `json:"text"`
	Author    User   `json:"author"`
	DateTime  string `json:"dateTime"`
}
