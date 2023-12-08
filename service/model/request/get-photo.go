package request

type GetPhotoRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
}

func (request *GetPhotoRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId)
}
