package request

type GetCommentsRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
	QueryParameters struct {
		StartDate string `key:"startDate"`
		StartID   string `key:"startID"`
	}
}

func (request *GetCommentsRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) &&
		isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartID == "" || isValidUUID(request.QueryParameters.StartID))
}
