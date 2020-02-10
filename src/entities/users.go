package entities

import "time"

type User struct {
	ID             uint   `gorm:"primary_key"`
	DisplayID      string `gorm:"unique;not null"`
	IsUnsubscribed bool   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
