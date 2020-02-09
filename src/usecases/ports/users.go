package ports

import (
	"time"

	"github.com/GraphQLSample/src/entities"
)

type UsersOutputPort struct {
	Users []entities.User `json:"users"`
}

type UserOutputPort struct {
	ID             uint      `json:"id"`
	IsUnsubscribed bool      `json:"is_unsubscribed"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserInputPort struct {
	UserID uint `json:"user_id"`
}
