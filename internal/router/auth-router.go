package router

import (
	"github.com/gofiber/fiber/v2"
	"gogram/internal/app/auth"
)

func AuthRouter(r fiber.Router) {
	authHandler := auth.NewAuthHandler()

	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
}

func UserRouter(r fiber.Router) {
	authHandler := auth.NewAuthHandler()

	r.Get("/profile", authHandler.Profile)
}
