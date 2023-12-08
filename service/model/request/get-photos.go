package request

type GetPhotosRequest struct {
	QueryParameters struct {
		UserId    string `key:"userId"`
		StartDate string `key:"startDate"`
		StartId   string `key:"startId"`
	}
}

func (request *GetPhotosRequest) IsValid() bool {
	return isValidUUID(request.QueryParameters.UserId) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
