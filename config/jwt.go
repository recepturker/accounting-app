package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID    uint `json:"user_id"`
	PartnerID uint `json:"partner_id"`
	jwt.RegisteredClaims
}

// JWT Token oluştur
func GenerateToken(userID uint, partnerID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 1 gün geçerli

	claims := &Claims{
		UserID:    userID,
		PartnerID: partnerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT Token doğrula ve kullanıcı bilgilerini al
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("geçersiz token")
	}

	return claims, nil
}
