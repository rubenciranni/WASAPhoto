package request

type GetPhotosRequest struct {
	QueryParameters struct {
		UserID    string `key:"userID"`
		StartDate string `key:"startDate"`
		StartId   string `key:"startId"`
	}
}

func (request *GetPhotosRequest) IsValid() bool {
	return isValidUUID(request.QueryParameters.UserID) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartId == "" || isValidUUID(request.QueryParameters.StartId))
}
