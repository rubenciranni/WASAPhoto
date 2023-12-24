package schema

type UserProfile struct {
	UserId            string `json:"userId"`
	Username          string `json:"username"`
	NumberOfPhotos    int    `json:"numberOfPhotos"`
	NumberOfFollowers int    `json:"numberOfFollowers"`
	NumberOfFollowing int    `json:"numberOfFollowing"`
	IsFollowed        bool   `json:"isFollowed"`
}
