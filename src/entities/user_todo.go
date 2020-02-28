package entities

import "time"

type UserTodo struct {
	ID        uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"not null"`
	Text      string `gorm:"not null"`
	Done      bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
