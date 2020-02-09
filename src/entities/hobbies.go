package entities

import "time"

type Hobby struct {
	ID        uint
	Name      string
	IsDelete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
