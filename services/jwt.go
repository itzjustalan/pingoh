package services

import (
	"errors"
	"pingoh/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UID    string        `json:"uid"`
	Role   string        `json:"role"`
	Access db.UserAccess `json:"access"`
	jwt.RegisteredClaims
}

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var issuer = "pingoh"
var secret = "secret"

func NewJwtTokens(uid string, role string, access db.UserAccess) (JwtTokens, error) {
	var tokens JwtTokens
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		UID:    uid,
		Role:   role,
		Access: access,
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
		UID:    uid,
		Role:   role,
		Access: access,
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
