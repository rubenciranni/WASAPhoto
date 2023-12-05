package schema

type CommentList struct {
	LastDate string    `json:"lastDate"`
	LastId   string    `json:"lastId"`
	Records  []Comment `json:"records"`
}
