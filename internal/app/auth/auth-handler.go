package auth

import (
	"github.com/gofiber/fiber/v2"
	auth_requests2 "gogram/internal/app/auth/auth-requests"
	"gogram/pkg/helpers"
	"gogram/pkg/validation"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: NewAuthService(),
	}
}

func (ah *AuthHandler) Login(ctx *fiber.Ctx) error {
	req := new(auth_requests2.LoginRequest)

	if err := validation.Validate[auth_requests2.LoginRequest](ctx, req); err != nil {
		return err
	}

	tokenJWT, usr, err := ah.service.Login(req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := &AuthResponse{
		Token: tokenJWT,
		User:  usr,
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Login success",
		data,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ah *AuthHandler) Register(ctx *fiber.Ctx) error {
	req := new(auth_requests2.RegisterRequest)

	if err := validation.Validate[auth_requests2.RegisterRequest](ctx, req); err != nil {
		return err
	}

	tokenJWT, usr, err := ah.service.Register(req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := &AuthResponse{
		Token: tokenJWT,
		User:  usr,
	}

	res := helpers.NewResponseHelper(
		fiber.StatusCreated,
		"Register success",
		data,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ah *AuthHandler) Profile(ctx *fiber.Ctx) error {
	userClaims, err := User(ctx)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Get User Profile",
		userClaims,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}
