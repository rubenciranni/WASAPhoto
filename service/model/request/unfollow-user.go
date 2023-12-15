package request

type UnfollowUserRequest struct {
	PathParameters struct {
		UserID string `key:"photoID"`
	}
}

func (request *UnfollowUserRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.UserID)
}
