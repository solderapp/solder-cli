package solder

import (
	"time"
)

// Message represents a standard response.
type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Build represents a build API response.
type Build struct {
	ID          int64     `json:"id"`
	Pack        string    `json:"pack"`
	PackID      int64     `json:"pack_id"`
	Minecraft   string    `json:"minecraft"`
	MinecraftID int64     `json:"minecraft_id"`
	Forge       string    `json:"forge"`
	ForgeID     int64     `json:"forge_id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	MinJava     string    `json:"min_java"`
	MinMemory   string    `json:"min_memory"`
	Published   bool      `json:"published"`
	Private     bool      `json:"private"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Client represents a client API response.
type Client struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Value     string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Forge represents a forge API response.
type Forge struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Key represents a key API response.
type Key struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Value     string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Minecraft represents a minecraft API response.
type Minecraft struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Mod represents a mod API response.
type Mod struct {
	ID          int64     `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Pack represents a pack API response.
type Pack struct {
	ID            int64     `json:"id"`
	Slug          string    `json:"slug"`
	Name          string    `json:"name"`
	Icon          string    `json:"icon"`
	Logo          string    `json:"logo"`
	Background    string    `json:"background"`
	RecommendedID int64     `json:"recommended_id"`
	Recommended   string    `json:"recommended"`
	LatestID      int64     `json:"latest_id"`
	Latest        string    `json:"latest"`
	Website       string    `json:"website"`
	Hidden        bool      `json:"hidden"`
	Private       bool      `json:"private"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// User represents a user API response.
type User struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Version represents a version API response.
type Version struct {
	ID        int64     `json:"id"`
	Mod       string    `json:"mod"`
	ModID     int64     `json:"mod_id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Profile represents a profile API response.
type Profile struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
