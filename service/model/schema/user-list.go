package schema

type UserList struct {
	PreviousOffset int
	NextOffset     int
	HasPrevious    bool
	HasNext        bool
	records        []User
}
