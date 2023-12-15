package request

type UnbanUserRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
}

func (request *UnbanUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
