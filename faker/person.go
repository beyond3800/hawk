package faker

import "fmt"

var firstNames = []string{
	"Adam",
	"John",
	"James",
	"David",
	"Michael",
	"Sarah",
	"Grace",
	"Mary",
	"Faith",
	"Esther",
}

var lastNames = []string{
	"Smith",
	"Johnson",
	"Brown",
	"Williams",
	"Miller",
	"Amusa",
	"Adeyemi",
	"Okoro",
	"Bello",
	"Ahmed",
}

func FirstName() string {
	return random(firstNames)
}

func LastName() string {
	return random(lastNames)
}

func Name() string {
	return fmt.Sprintf("%s %s", FirstName(), LastName())
}