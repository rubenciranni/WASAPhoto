package request

type GetUserProfileRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
}

func (request *GetUserProfileRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
