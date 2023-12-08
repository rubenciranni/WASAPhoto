package request

import (
	"regexp"

	"github.com/gofrs/uuid"
)

func isValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}

var dateTimeRegexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)

func isValidDateTime(d string) bool {
	return dateTimeRegexp.MatchString(d)
}