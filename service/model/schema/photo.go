package schema

import (
	"time"

	"github.com/gofrs/uuid"
)

type Photo struct {
	PhotoId          uuid.UUID
	Author           User
	DateTime         time.Time
	Caption          string
	numberOfLikes    int
	numberOfComments int
}
