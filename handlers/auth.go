package handlers

import (
	"database/sql"

	"pingoh/db"
	"pingoh/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email string `json:"email" xml:"email" form:"email"`
	Passw string `json:"passw" xml:"passw" form:"passw"`
}

type AuthenticatedUser struct {
	db.User
	services.JwtTokens
}

func Signup(creds *AuthCredentials) (AuthenticatedUser, error) {
	var u AuthenticatedUser
	if len(creds.Passw) > 50 {
		return u, fiber.NewError(
			fiber.ErrBadRequest.Code,
			"password must be less than 50 characters!",
		)
	}
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Passw), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	_, err = db.CreateUser(&db.User{Email: creds.Email, PwHash: string(hash)})
	if err != nil {
		return u, err
	}
	user, err := db.FindUserByEmail(creds.Email)
	if err != nil {
		return u, err
	}
	tokens, err := services.NewJwtTokens(user.ID, string(user.Role), user.Access)
	if err != nil {
		return u, err
	}
	u.User = user
	u.JwtTokens = tokens
	return u, nil
}

func Signin(creds *AuthCredentials) (AuthenticatedUser, error) {
	var u AuthenticatedUser
	user, err := db.FindUserByEmail(creds.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, fiber.NewError(
				fiber.ErrBadRequest.Code, "user not found")
		}
		return u, err
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(user.PwHash), []byte(creds.Passw)) != nil {
		return u, fiber.NewError(
			fiber.ErrBadRequest.Code, "wrong password")
	}
	tokens, err := services.NewJwtTokens(user.ID, string(user.Role), user.Access)
	if err != nil {
		return u, err
	}
	u.User = user
	u.JwtTokens = tokens
	return u, nil
}

func RefreshTokens(token string) (services.JwtTokens, error) {
	var tokens services.JwtTokens
	claims, err := services.ValidateToken(token)
	if err != nil {
		return tokens, err
	}
	tokens, err = services.NewJwtTokens(claims.ID, claims.Role, claims.Access)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}
