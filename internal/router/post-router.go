package router

import (
	"github.com/gofiber/fiber/v2"
	"gogram/internal/app/post"
)

func PostRouter(r fiber.Router) {
	postHandler := post.NewPostHandler()

	r.Get("/posts", postHandler.Index)
	r.Post("/posts", postHandler.Create)
	r.Get("/posts/:id", postHandler.Show)
	r.Put("/posts/:id", postHandler.Update)
	r.Delete("/posts/:id", postHandler.Delete)
	r.Get("/external-posts", postHandler.GetFromJSONPlaceholder)
}
