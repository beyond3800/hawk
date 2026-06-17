package lib

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string, error) {
	hashedPwd,err:= bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hashedPwd),nil
}

func PasswordVerify(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}