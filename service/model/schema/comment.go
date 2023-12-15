package schema

type Comment struct {
	CommentID string `json:"commentID"`
	Text      string `json:"text"`
	Author    User   `json:"author"`
	DateTime  string `json:"dateTime"`
}
