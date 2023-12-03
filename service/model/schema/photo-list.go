package schema

type PhotoList struct {
	PreviousOffset int
	NextOffset     int
	HasPrevious    bool
	HasNext        bool
	records        []Photo
}
