package faker

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func random(list []string) string {
	return list[rnd.Intn(len(list))]
}


// faker/
//     faker.go
//     person.go
//     internet.go
//     text.go
//     number.go
//     security.go
//     time.go