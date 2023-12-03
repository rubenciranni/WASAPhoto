package schema

type CommentList struct {
	PreviousOffset int
	NextOffset     int
	HasPrevious    bool
	HasNext        bool
	records        []Comment
}
