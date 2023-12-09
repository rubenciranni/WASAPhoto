package request

type GetFollowingRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetFollowingRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
