package request

type SetMyUserNameRequest struct {
	NewUsername string `json:"newUsername"`
}

func (request *SetMyUserNameRequest) IsValid() bool {
	return usernameRegexp.MatchString(request.NewUsername)
}
