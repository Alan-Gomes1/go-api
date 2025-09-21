package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
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

func VerifyTokenMIddleware(c *gin.Context) {
	secret := os.Getenv(SECRET_KEY_JWT)
	tokenValue := c.Request.Header.Get("Authorization")

	newToken := RemoveBearerPrefix(tokenValue)
	token, err := jwt.Parse(newToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		restErr := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(restErr.Code, restErr)
		c.Abort()
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		restErr := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(restErr.Code, restErr)
		c.Abort()
		return
	}
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix("Bearer ", token) {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}
