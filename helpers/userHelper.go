package helpers

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const SECRET = "expenseSecret"

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(hashedPassword)
}

func ComparePasswords(hashPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GenerateJwtToken(userId string, admin bool) (string, error) {
	//Create token
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//Generate encoded token and send it as response
	return token.SignedString([]byte(SECRET))
}

func ValidateJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["userId"])
		return claims, nil
	} else {
		// fmt.Println(err)
		return nil, err
	}
}
