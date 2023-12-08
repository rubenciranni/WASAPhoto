package request

type GetPhotosRequest struct {
	QueryParameters struct {
		UserId   string
		LastDate string
		LastId   string
	}
}

func (request *GetPhotosRequest) IsValid() bool {
	return isValidUUID(request.QueryParameters.UserId) &&
		isValidDateTime(request.QueryParameters.LastDate) &&
		(request.QueryParameters.LastId == "" || isValidUUID(request.QueryParameters.LastId))
}
