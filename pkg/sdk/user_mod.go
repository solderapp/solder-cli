package sdk

// UserMod represents a user mod API response.
type UserMod struct {
	User *User  `json:"user,omitempty"`
	Mod  *Mod   `json:"mod,omitempty"`
	Perm string `json:"perm,omitempty"`
}
