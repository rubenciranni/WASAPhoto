package request

type SearchUserRequest struct {
	QueryParameters struct {
		Username     string `key:"username"`
		IsExactMatch bool   `key:"isExactMatch"`
		StartId      string `key:"startId"`
	}
}

func (request *SearchUserRequest) IsValid() bool {
	return request.QueryParameters.StartId == "" ||
		isValidUUID(request.QueryParameters.StartId)
}
