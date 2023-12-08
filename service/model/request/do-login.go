package request

type DoLoginRequest struct {
	Username string `json:"username"`
}

func (request *DoLoginRequest) IsValid() bool {
	return isValidUsername(request.Username)
}
