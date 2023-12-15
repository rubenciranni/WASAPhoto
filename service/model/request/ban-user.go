package request

type BanUserRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
}

func (request *BanUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
