package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gogram/config"
	"gogram/internal/app/user"
	"strings"
	"time"
)

type UserClaims struct {
	User user.User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateJWT(usr user.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &UserClaims{
		User: usr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString(config.JwtSecret)))
}

func ParseJWT(tokenStr string) (*UserClaims, error) {
	claims := &UserClaims{}

	jwToken, err := jwt.ParseWithClaims(
		strings.TrimPrefix(tokenStr, "Bearer "),
		claims,
		func(tkn *jwt.Token) (any, error) {
			return []byte(viper.GetString(config.JwtSecret)), nil
		})

	if err != nil {
		return nil, err
	}

	if !jwToken.Valid {
		return nil, errors.New("invalid Token")
	}

	return claims, nil
}
