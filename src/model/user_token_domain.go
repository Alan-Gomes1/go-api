package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
)

const SECRET_KEY_JWT = "SECRET_KEY_JWT"

func (u *userDomain) GenerateToken() (string, *rest_err.Errors) {
	secret := os.Getenv(SECRET_KEY_JWT)
	claims := jwt.MapClaims{
		"id":    u.id,
		"email": u.email,
		"name":  u.name,
		"age":   u.age,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		error_msg := fmt.Sprintf("Error trying to generate token: %s", err)
		return "", rest_err.NewInternalServerError(error_msg)
	}

	return signedToken, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.Errors) {
	secret := os.Getenv(SECRET_KEY_JWT)
	if tokenValue == "" {
		return nil, rest_err.NewBadRequestError("invalid token")
	}

	newToken := RemoveBearerPrefix(tokenValue)
	token, err := jwt.Parse(newToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	user := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}
	return user, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix("Bearer ", token) {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}
