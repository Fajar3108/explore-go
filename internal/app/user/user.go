package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"id" gorm:"primaryKey;autoIncrement:true;not null"`
	Username  string       `json:"username" gorm:"not null"`
	Password  string       `json:"-" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at" gorm:"not null"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
