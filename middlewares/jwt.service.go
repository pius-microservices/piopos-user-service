package middlewares

import (
	"fmt"
	"time"

	envConfig "github.com/pius-microservices/piopos-user-service/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(envConfig.LoadConfig().JwtSecret)

type Claims struct {
	UserId string
	jwt.RegisteredClaims
}

func NewToken(id string, expiresIn time.Duration) *Claims {
	return &Claims{
		UserId:           id,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn))},
	}
}

func (claim *Claims) CreateToken() (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return accessToken.SignedString(jwtSecret)
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}