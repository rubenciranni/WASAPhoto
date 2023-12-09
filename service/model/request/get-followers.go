package request

type GetFollowersRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetFollowersRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
