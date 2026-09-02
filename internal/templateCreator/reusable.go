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
		newstring,_:=strings.CutSuffix(name,"y",)
		newName = newstring+"ies"
	case strings.HasSuffix(name,"s"):
		newName = name
	default:
		newName = name+"s"
	}
	return newName
}