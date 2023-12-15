package request

type GetCommentsRequest struct {
	PathParameters struct {
		PhotoId string `key:"photoId"`
	}
	QueryParameters struct {
		StartDate string `key:"startDate"`
		StartId   string `key:"startId"`
	}
}

func (request *GetCommentsRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoId) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
