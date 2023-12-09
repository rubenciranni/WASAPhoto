package request

type LikePhotoRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
}

func (request *LikePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId)
}
