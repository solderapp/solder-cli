package sdk

import (
	"time"
)

// Key represents a key API response.
type Key struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Value     string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Key) String() string {
	return s.Name
}
