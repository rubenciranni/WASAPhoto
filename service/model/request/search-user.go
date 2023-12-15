package request

type SearchUserRequest struct {
	QueryParameters struct {
		Username string `key:"username"`
		StartId  string `key:"startId"`
	}
}

func (request *SearchUserRequest) IsValid() bool {
	return request.QueryParameters.StartId == "" ||
		isValidUUID(request.QueryParameters.StartId)
}
