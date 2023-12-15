package request

type GetLikesRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetLikesRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
