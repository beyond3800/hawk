package lib

import (
	"strings"
)

func IsValid(input string) bool {
	word := strings.Split(input, "_")
	if len(word) != 3{
		return false
	}
	return word[0] == "create"
}