package request

type GetPhotoRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
}

func (request *GetPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID)
}
