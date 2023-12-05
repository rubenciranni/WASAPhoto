package schema

type Photo struct {
	PhotoId          string `json:"photoId"`
	Author           User   `json:"author"`
	DateTime         string `json:"dateTime"`
	Caption          string `json:"caption"`
	NumberOfLikes    int    `json:"numberOfLikes"`
	NumberOfComments int    `json:"numberOfComments"`
}
