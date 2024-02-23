package utils

import (
	"fmt"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Jwt interface{
	GenerateToken(userID uuid.UUID, username string) (string, error)
	VerifyToken(tokenString string) (jwt.Claims, error)
}

var secretKey = []byte(getSecretkey())

func getSecretkey() string {
	secret := os.Getenv("SECRET_KEY")

	if secret == ""{
		secret = "secretkey"
	}

	return secret
}

func GenerateToken(userID uuid.UUID, username string) (string, error){
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour*1).Unix() // Token valid for one hour

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.Claims, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Invalid signing method")
        }

        return secretKey, nil
    })

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid Token")
}