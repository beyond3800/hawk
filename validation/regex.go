package validation

import "regexp"

var stringRegex = regexp.MustCompile(
	`^[A-Za-z_][A-Za-z0-9_]*$`,
)

var emailRegex = regexp.MustCompile(
	`^[^\s@]+@[^\s@]+\.[^\s@]+$`,
)