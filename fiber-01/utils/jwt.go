package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecrect = []byte(getJWTSecrect())

func getJWTSecrect() string {
	secrect := os.Getenv("JWT_SECRECT")
	if secrect == "" {
		secrect = "dev-secrect"
	}
	return secrect
}

type JWTClaims struct {
	AccountID uint `json:"account_id"`
	jwt.RegisteredClaims
}

func GenerateToken(accountID uint) (string, error){
	claims := JWTClaims {
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return  token.SignedString(jwtSecrect)
}


func ParseToken(tokenStr string) (uint, error){

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecrect, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.AccountID, nil

}

