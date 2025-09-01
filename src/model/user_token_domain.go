package model

import (
	"fmt"
	"os"
	"time"

	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
)

const SECRET_KEY_JWT = "SECRET_KEY_JWT"

func (u *userDomain) GenerateToken() (string, *rest_err.Errors) {
	secret := os.Getenv(SECRET_KEY_JWT)
	claims := jwt.MapClaims{
		"id": u.id,
		"email": u.email,
		"name": u.name,
		"age": u.age,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		error_msg := fmt.Sprintf("Error trying to generate token: %s", err)
		return "", rest_err.NewInternalServerError(error_msg)
	}

	return signedToken, nil
}
