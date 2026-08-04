package faker

import (
	"fmt"
	"strings"
)

var domains = []string{
	"gmail.com",
	"yahoo.com",
	"hotmail.com",
	"example.com",
}

func Username() string {
	name := strings.ToLower(FirstName())
	return fmt.Sprintf("%s%d", name, rnd.Intn(9999))
}

func Email() string {
	return fmt.Sprintf("%s@%s", Username(), random(domains))
}