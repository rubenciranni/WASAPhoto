package schema

type UserList struct {
	LastDate string `json:"lastDate"`
	LastId   string `json:"lastId"`
	Records  []User `json:"records"`
}
