package kleister

import (
	"time"
)

// Mod represents a mod API response.
type Mod struct {
	ID          int64      `json:"id"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	Website     string     `json:"website"`
	Donate      string     `json:"donate"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Users       []*User    `json:"users,omitempty"`
	UserMods    []*UserMod `json:"user_mods,omitempty"`
	Teams       []*Team    `json:"teams,omitempty"`
	TeamMods    []*TeamMod `json:"team_mods,omitempty"`
}

func (s *Mod) String() string {
	return s.Name
}
