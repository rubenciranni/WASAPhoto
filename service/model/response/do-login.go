package response

import "github.com/gofrs/uuid"

type DoLoginResponse struct {
	UserId uuid.UUID
}
