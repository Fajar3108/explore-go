package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gogram/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	if db != nil {
		return db
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString(config.DbUser),
		viper.GetString(config.DbPassword),
		viper.GetString(config.DbHost),
		viper.GetString(config.DbPort),
		viper.GetString(config.DbDatabase),
	)

	var err error

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	return db
}
