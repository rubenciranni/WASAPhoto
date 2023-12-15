package request

type UnlikePhotoRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
}

func (request *UnlikePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID)
}
