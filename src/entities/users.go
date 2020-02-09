package entities

import "time"

type User struct {
	ID             uint
	IsUnsubscribed bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
