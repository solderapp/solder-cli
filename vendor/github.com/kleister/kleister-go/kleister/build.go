package kleister

import (
	"time"
)

// Build represents a build API response.
type Build struct {
	ID            int64           `json:"id"`
	Pack          *Pack           `json:"pack,omitempty"`
	PackID        int64           `json:"pack_id"`
	Minecraft     *Minecraft      `json:"minecraft,omitempty"`
	MinecraftID   int64           `json:"minecraft_id"`
	Forge         *Forge          `json:"forge,omitempty"`
	ForgeID       int64           `json:"forge_id"`
	Slug          string          `json:"slug"`
	Name          string          `json:"name"`
	MinJava       string          `json:"min_java"`
	MinMemory     string          `json:"min_memory"`
	Published     bool            `json:"published"`
	Private       bool            `json:"private"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Versions      []*Version      `json:"versions,omitempty"`
	BuildVersions []*BuildVersion `json:"build_versions,omitempty"`
}

func (s *Build) String() string {
	return s.Name
}
