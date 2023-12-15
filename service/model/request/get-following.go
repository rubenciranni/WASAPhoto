package request

type GetFollowingRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
	QueryParameters struct {
		StartID string `key:"startID"`
	}
}

func (request *GetFollowingRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
