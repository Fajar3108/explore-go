package router

import (
	"github.com/gofiber/fiber/v2"
	errorhandler "gogram/pkg/error-handler"
	"gogram/pkg/middlewares"
)

func SetupRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorhandler.GlobalErrorHandler,
	})

	api := app.Group("api")
	AuthRouter(api)

	api.Use(middlewares.JWTMiddleware())
	UserRouter(api)
	PostRouter(api)

	return app
}
