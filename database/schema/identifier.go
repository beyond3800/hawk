package database

import "regexp"

var identifierRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func validIdentifier(name string) bool {
	return identifierRegex.MatchString(name)
}