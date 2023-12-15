package request

type FollowUserRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
}

func (request *FollowUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
