package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type AuthUserClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}

	return header
}
func ExtractUserIDFromToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}
	fmt.Println(claims.Id, "id is ")

	return claims.Id, claims.Email, nil

}
