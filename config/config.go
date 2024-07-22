package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	JwtSecret  = "JWT_SECRET"
	DbHost     = "DB_HOST"
	DbPort     = "DB_PORT"
	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
	DbDatabase = "DB_DATABASE"
)

func InitConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()
}
