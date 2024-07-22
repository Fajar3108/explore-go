package post

import (
	"database/sql"
	"gogram/internal/app/user"
	"time"
)

type Post struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Thumbnail sql.NullString `json:"thumbnail"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UserID    uint           `json:"-"`
	User      user.User      `json:"user" gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}
