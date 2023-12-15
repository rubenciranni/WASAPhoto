package request

type GetLikesRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
	QueryParameters struct {
		StartId string `key:"startId"`
	}
}

func (request *GetLikesRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
