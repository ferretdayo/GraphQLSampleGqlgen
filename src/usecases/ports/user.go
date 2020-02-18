package ports

import (
	"time"
)

type UserOutputPort struct {
	ID             uint      `json:"id"`
	DisplayID      string    `json:"display_id"`
	IsUnsubscribed bool      `json:"is_unsubscribed"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Hobby		   *MasterOutputPort `json:"hobby"`
}

type UserInputPort struct {
	UserID uint `json:"user_id"`
}
