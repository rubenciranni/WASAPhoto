package schema

import (
	"github.com/gofrs/uuid"
)

type UserProfile struct {
	UserId            uuid.UUID
	Username          string
	NumberOfPhotos    int
	NumberOfFollowers int
	NumberOfFollowing int
}
