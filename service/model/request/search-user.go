package request

type SearchUserRequest struct {
	QueryParameters struct {
		Username string `key:"username"`
		StartID  string `key:"startID"`
	}
}

func (request *SearchUserRequest) IsValid() bool {
	return request.QueryParameters.StartID == "" ||
		isValidUUID(request.QueryParameters.StartID)
}
