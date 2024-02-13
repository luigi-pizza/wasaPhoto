package requests

import (
	"regexp"

	"github.com/luigi-pizza/wasaPhoto/service/filesystem"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z][\\.]{0,1}(?:[\\w][\\.]{0,1})*[\\w]$`)
var usernameSearchRegex = regexp.MustCompile(`^[a-zA-Z][\\.]{0,1}(?:[\\w][\\.]{0,1})*$`)

func check_username(username string) bool {
	return 5 <= len(username) && len(username) <= 25 && usernameRegex.MatchString(username)
}

func check_usernameSearch(username string) bool {
	return 3 <= len(username) && len(username) <= 25 && usernameSearchRegex.MatchString(username)
}

func check_commentText(text string) bool {
	return len(text) != 0 && len(text) < filesystem.MaxTextLength
}
