package schema

import (
	"github.com/gofrs/uuid"
)

type User struct {
	UserId   uuid.UUID
	Username string
}
