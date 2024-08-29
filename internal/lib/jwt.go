package lib

import (
	"auth/service/internal/domain/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWTToken(u model.User, secret string, duraction time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = u.Id
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(duraction).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
