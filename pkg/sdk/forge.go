package sdk

import (
	"time"
)

// Forge represents a forge API response.
type Forge struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Forge) String() string {
	return s.Version
}
