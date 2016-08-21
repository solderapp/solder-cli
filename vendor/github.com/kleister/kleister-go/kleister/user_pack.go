package kleister

// UserPack represents a user pack API response.
type UserPack struct {
	User *User  `json:"user,omitempty"`
	Pack *Pack  `json:"pack,omitempty"`
	Perm string `json:"perm,omitempty"`
}
