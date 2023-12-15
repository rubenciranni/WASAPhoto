package request

type BanUserRequest struct {
	PathParameters struct {
		UserId string `key:"userId"`
	}
}

func (request *BanUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserId)
}
