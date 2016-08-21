package kleister

// TeamMod represents a team mod API response.
type TeamMod struct {
	Team *Team  `json:"team,omitempty"`
	Mod  *Mod   `json:"mod,omitempty"`
	Perm string `json:"perm,omitempty"`
}
