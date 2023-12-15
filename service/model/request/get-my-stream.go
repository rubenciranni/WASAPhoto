package request

type GetMyStreamRequest struct {
	QueryParameters struct {
		StartDate string `key:"startDate"`
		StartID   string `key:"startID"`
	}
}

func (request *GetMyStreamRequest) IsValid() bool {
	return isValidDateTime(request.QueryParameters.StartDate) &&
		(request.QueryParameters.StartID == "" || isValidUUID(request.QueryParameters.StartID))
}
