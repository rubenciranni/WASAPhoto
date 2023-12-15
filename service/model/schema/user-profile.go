package schema

type UserProfile struct {
	UserID            string `json:"userID"`
	Username          string `json:"username"`
	NumberOfPhotos    int    `json:"numberOfPhotos"`
	NumberOfFollowers int    `json:"numberOfFollowers"`
	NumberOfFollowing int    `json:"numberOfFollowing"`
}
