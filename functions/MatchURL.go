package functions

import (
	"regexp"
)

func MatchURL(url string) bool {
	regex := `^((http(s)?:\/\/)?(www\.)?[a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{1,24}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*))$`

	match, err := regexp.Match(regex, []byte(url))
	if !match || err != nil {
		return false
	}
	return true
}
