package request

import "regexp"

var usernameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)

type DoLoginRequest struct {
	Username string `json:"username"`
}

func (request *DoLoginRequest) IsValid() bool {
	return usernameRegexp.MatchString(request.Username)
}
