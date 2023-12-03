package schema

type UserProfile struct {
	UserId            UUID
	Username          string
	NumberOfPhotos    int
	NumberOfFollowers int
	NumberOfFollowing int
}
