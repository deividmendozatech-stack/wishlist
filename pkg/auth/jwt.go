package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey used for signing JWT tokens.
var SecretKey = []byte("clave-secreta")

// GenerateToken creates a signed JWT valid for 24h containing the user ID.
//
// Params:
//
//	userID - identifier of the authenticated user
//
// Returns:
//
//	string - the signed JWT
//	error  - if signing fails
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
