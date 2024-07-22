package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	auth_requests2 "gogram/internal/app/auth/auth-requests"
	"gogram/internal/app/user"
	"gogram/internal/database"
	"gogram/pkg/mail"
	"gogram/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		db: database.InitDB(),
	}
}

func (as *AuthService) Login(request *auth_requests2.LoginRequest) (string, *user.User, error) {
	usr := new(user.User)

	result := as.db.Where("username = ?", request.Username).First(usr)

	if result.Error != nil {
		return "", nil, fiber.NewError(fiber.StatusUnauthorized, "Username not registered")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(request.Password)); err != nil {
		return "", nil, fiber.NewError(fiber.StatusUnauthorized, "Password is incorrect")
	}

	jwToken, err := token.GenerateJWT(*usr)

	if err != nil {
		return "", nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return jwToken, usr, nil
}

func (as *AuthService) Register(request *auth_requests2.RegisterRequest) (string, *user.User, error) {
	usr := &user.User{
		Username: request.Username,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	usr.Password = string(hashedPassword)

	if err := mail.SendMail(
		"maulanafajaribrahim@gmail.com",
		"Go Gram Register Success",
		fmt.Sprintf("<h1>Register Success</h1><p>Thank you for registering <b>%s</b></p>", usr.Username),
	); err != nil {
		return "", nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	result := as.db.Create(usr)

	if result.Error != nil {
		return "", nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	jwToken, err := token.GenerateJWT(*usr)

	if err != nil {
		return "", nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return jwToken, usr, nil
}
