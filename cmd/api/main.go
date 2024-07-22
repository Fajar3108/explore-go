package main

import (
	"gogram/config"
	"gogram/internal/router"
)

func main() {
	config.InitConfig()

	app := router.SetupRouter()

	app.Static("public", "./public")
	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
