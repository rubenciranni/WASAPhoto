package request

type FollowUserRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
}

func (request *FollowUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
