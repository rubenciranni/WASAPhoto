package request

type GetFollowingRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetFollowingRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
