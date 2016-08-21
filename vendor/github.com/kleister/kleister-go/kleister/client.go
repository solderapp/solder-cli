package kleister

import (
	"time"
)

// Client represents a client API response.
type Client struct {
	ID          int64         `json:"id"`
	Slug        string        `json:"slug"`
	Name        string        `json:"name"`
	Value       string        `json:"uuid"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Packs       []*Pack       `json:"packs,omitempty"`
	ClientPacks []*ClientPack `json:"client_packs,omitempty"`
}

func (s *Client) String() string {
	return s.Name
}
