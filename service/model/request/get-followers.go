package request

type GetFollowersRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
	QueryParameters struct {
		StartID string `key:"startID"`
	}
}

func (request *GetFollowersRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
