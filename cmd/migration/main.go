package main

import (
	"gogram/config"
	"gogram/internal/app/post"
	"gogram/internal/app/user"
	"gogram/internal/database"
)

func main() {
	config.InitConfig()

	db := database.InitDB()

	err := db.AutoMigrate(&post.Post{}, &user.User{})

	if err != nil {
		panic(err)
	}
}
