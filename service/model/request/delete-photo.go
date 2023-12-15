package request

type DeletePhotoRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
}

func (request *DeletePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID)
}
