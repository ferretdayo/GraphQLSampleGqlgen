package ports

import "time"

type MasterOutputPort struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
