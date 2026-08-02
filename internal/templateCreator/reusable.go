package templatecreator

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


func ToTitle(s string) string{
	caser := cases.Title(language.English)
	return caser.String(s)
}

func Pluralize(name string) string{
	var newName string
	switch {
	case strings.HasSuffix(name, "y"):
		newName = strings.Replace(newName,"y","ies",1)
	case strings.HasSuffix(name,"s"):
		newName = name
	default:
		newName = name+"s"
	}
	return newName
}