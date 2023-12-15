package request

type GetMyStreamRequest struct {
	QueryParameters struct {
		StartDate string `key:"startDate"`
		StartId   string `key:"startId"`
	}
}

func (request *GetMyStreamRequest) IsValid() bool {
	return isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
