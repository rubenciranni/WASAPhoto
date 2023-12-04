package schema

type Comment struct {
	CommentId string
	Text      string
	Author    User
	DateTime  string
}
