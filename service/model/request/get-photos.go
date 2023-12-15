package request

type GetPhotosRequest struct {
	QueryParameters struct {
		UserID    string `key:"userID"`
		StartDate string `key:"startDate"`
		StartID   string `key:"startID"`
	}
}

func (request *GetPhotosRequest) IsValid() bool {
	return isValidUUID(request.QueryParameters.UserID) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartID == "" || isValidUUID(request.QueryParameters.StartID))
}
