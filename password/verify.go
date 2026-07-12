package password


import "golang.org/x/crypto/bcrypt"


func Verify(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}