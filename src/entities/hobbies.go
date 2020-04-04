package entities

import "time"

type Hobby struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	IsDelete  bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
