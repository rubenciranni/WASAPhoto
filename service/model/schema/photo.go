package schema

import (
	"time"
)

type Photo struct {
	PhotoId          UUID
	Author           User
	DateTime         time.Time
	Caption          string
	numberOfLikes    int
	numberOfComments int
}
