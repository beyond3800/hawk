package util

import (
	"fmt"
	"math/rand"
)

func BOID(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b:= make([]byte,length)
	for character := range b {
		b[character] = characters[rand.Intn(len(characters))]
	}
	id:= fmt.Sprintf("ops%s",string(b))
	return id
}