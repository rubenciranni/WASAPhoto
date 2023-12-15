package request

type GetCommentsRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
	QueryParameters struct {
		StartDate string `key:"startDate"`
		StartId   string `key:"startId"`
	}
}

func (request *GetCommentsRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
