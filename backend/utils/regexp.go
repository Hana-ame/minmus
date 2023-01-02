package utils

import "regexp"

func IsValidUsername(username string) bool {
	// regexp.MustCompile(`\w+`)
	matched, err := regexp.MatchString(`\w+`, username)
	if err != nil {
		return false
	}
	return matched
}
