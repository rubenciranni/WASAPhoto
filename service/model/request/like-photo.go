package request

type LikePhotoRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
}

func (request *LikePhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID)
}
