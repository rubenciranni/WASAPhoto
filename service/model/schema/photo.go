package schema

type Photo struct {
	PhotoID          string `json:"photoID"`
	Author           User   `json:"author"`
	DateTime         string `json:"dateTime"`
	Caption          string `json:"caption"`
	NumberOfLikes    int    `json:"numberOfLikes"`
	NumberOfComments int    `json:"numberOfComments"`
}
