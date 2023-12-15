package request

type UnfollowUserRequest struct {
	PathParameters struct {
		UserId string `key:"photoId"`
	}
}

func (request *UnfollowUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
