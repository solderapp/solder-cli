package sdk

import (
	"time"
)

// Minecraft represents a minecraft API response.
type Minecraft struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Minecraft) String() string {
	return s.Version
}
