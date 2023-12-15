package schema

type UserList struct {
	LastID  string `json:"lastID"`
	Records []User `json:"records"`
}
