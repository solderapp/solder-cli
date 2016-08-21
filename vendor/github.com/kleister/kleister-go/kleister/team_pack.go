package kleister

// TeamPack represents a team pack API response.
type TeamPack struct {
	Team *Team  `json:"team,omitempty"`
	Pack *Pack  `json:"pack,omitempty"`
	Perm string `json:"perm,omitempty"`
}
