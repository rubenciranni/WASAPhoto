package schema

type CommentList struct {
	LastDate string    `json:"lastDate"`
	Records  []Comment `json:"records"`
}
