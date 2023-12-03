package schema

import (
	"time"

	"github.com/gofrs/uuid"
)

type Comment struct {
	CommentId uuid.UUID
	Text      string
	Author    User
	DateTime  time.Time
}
