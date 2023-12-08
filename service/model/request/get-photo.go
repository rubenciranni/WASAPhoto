package request

type GetPhotoRequest struct {
	PathParameters struct {
		PhotoId string
	}
}

func (request *GetPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId)
}
