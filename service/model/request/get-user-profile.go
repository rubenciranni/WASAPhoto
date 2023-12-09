package request

type GetUserProfileRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
}

func (request *GetUserProfileRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
