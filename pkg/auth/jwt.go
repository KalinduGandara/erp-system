package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secretKey      string
	expiryDuration time.Duration
}

func NewJWTService(secretKey string, expiryStr string) (*JWTService, error) {
	duration, err := time.ParseDuration(expiryStr)
	if err != nil {
		return nil, err
	}

	return &JWTService{
		secretKey:      secretKey,
		expiryDuration: duration,
	}, nil
}

func (s *JWTService) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(s.expiryDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", jwt.ErrSignatureInvalid
}
