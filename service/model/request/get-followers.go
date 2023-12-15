package request

type GetFollowersRequest struct {
	PathParameters struct {
		UserID string `key:"userID"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetFollowersRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
