package schema

type Photo struct {
	PhotoId          string
	Author           User
	DateTime         string
	Caption          string
	NumberOfLikes    int
	NumberOfComments int
}
