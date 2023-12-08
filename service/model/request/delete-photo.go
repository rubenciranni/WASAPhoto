package request

type DeletePhotoRequest struct {
	PathParameters struct {
		PhotoId string
	}
}

func (request *DeletePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId)
}
