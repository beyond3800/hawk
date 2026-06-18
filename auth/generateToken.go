package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beyond3800/hawk/types"
	_"github.com/beyond3800/hawk/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey = []byte(os.Getenv("LOG_KEY"))


func Auth(user_id string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	jti := `util.UniqueId("expired_tokens")`
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":user_id,
		"jti": jti,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err:= token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString,nil
	// fmt.Print(jwtKey)
}


func ValidateToken(tokenStr string) (types.Token, error) {
	var tokenDetails types.Token
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			fmt.Println(ok)
			return nil,fmt.Errorf("Invalid token")
		}
        return jwtKey, nil
    })
	
    if err != nil || !token.Valid {
        return tokenDetails, fmt.Errorf("Invalid token")
    }
    claims,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return tokenDetails, fmt.Errorf("Invalid token")
	}
    uid,ok := claims["user_id"].(string)
	if !ok{
		return tokenDetails, fmt.Errorf("Invalid user_id")
	}
    jti,ok := claims["jti"].(string)
	if !ok{
		return tokenDetails, fmt.Errorf("Invalid jti")
	}
	exp,ok := claims["exp"].(float64)
	if !ok{

		return tokenDetails, fmt.Errorf("Invalid exp")
	}
	expTime := time.Unix(int64(exp),0)
	tokenDetails.Jti = jti
	tokenDetails.User_id = uid
	tokenDetails.Exp = expTime
    return tokenDetails, nil
}