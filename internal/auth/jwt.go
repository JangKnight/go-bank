package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"example.com/go-bank/internal/domain"
)

const secret = "SUPASAFESECRETSECRET"

type MyClaims struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    jwt.RegisteredClaims
}

func GenerateToken(user domain.User) (string, error) {
    userid := strconv.Itoa(user.ID)
    claims := MyClaims{
        Username: user.Name,
        Email:    user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            Issuer:    "go-bank-server",
            Subject:   userid,
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}