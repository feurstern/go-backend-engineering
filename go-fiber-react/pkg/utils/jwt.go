package utils

import (
	"go-fiber-react/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = []byte(config.Secret)

func GenerateJWT(UserID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SecretKey)
}
