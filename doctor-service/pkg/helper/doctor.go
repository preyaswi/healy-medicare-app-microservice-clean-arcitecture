package helper

import (
	"doctor-service/pkg/models"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashPassword)
	return hash, nil
}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	userId := strconv.Itoa(userID)
	claims := &AuthUserClaims{
		Id:    userId,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GenerateAccessToken(user models.DoctorDetail) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(int(user.Id), user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(user models.DoctorDetail) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := GenerateTokenUsers(int(user.Id), user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
