package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JwtTokens struct {
	AccessToken  string
	RefreshToken string
}

var issuer = "pingoh"
var secret = "secret"

func NewJwtTokens(id int, role string) (JwtTokens, error) {
	var tokens JwtTokens
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}).SignedString([]byte(secret))
	if err != nil {
		return tokens, err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	}).SignedString([]byte(secret))
	if err != nil {
		return tokens, err
	}

	return JwtTokens{accessToken, refreshToken}, nil
}

func ParseTokenClaims(tokenString string) (*UserClaims, error) {
	var claims *UserClaims
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, err
	} else {
		return claims, errors.New("bad claims")
	}
}
