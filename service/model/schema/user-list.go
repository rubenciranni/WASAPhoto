package schema

type UserList struct {
	LastId  string `json:"lastId"`
	Records []User `json:"records"`
}
