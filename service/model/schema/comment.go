package schema

import (
	"time"
)

type Comment struct {
	CommentId UUID
	Text      string
	Author    User
	DateTime  time.Time
}
