package request

type UnlikePhotoRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
}

func (request *UnlikePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId)
}
