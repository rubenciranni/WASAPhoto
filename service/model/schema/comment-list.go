package schema

type CommentList struct {
	LastDate string    `json:"lastDate"`
	LastID   string    `json:"lastID"`
	Records  []Comment `json:"records"`
}
