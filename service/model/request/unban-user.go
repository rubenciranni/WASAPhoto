package request

type UnbanUserRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
}

func (request *UnbanUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
