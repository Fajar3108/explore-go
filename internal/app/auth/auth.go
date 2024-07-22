package auth

import (
	"github.com/gofiber/fiber/v2"
	"gogram/internal/app/user"
	"gogram/pkg/token"
)

type AuthResponse struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}

func User(ctx *fiber.Ctx) (*token.UserClaims, error) {
	authHeader := ctx.GetReqHeaders()["Authorization"][0]

	userClaims, err := token.ParseJWT(authHeader)

	return userClaims, err
}
