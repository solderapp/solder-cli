package kleister

// Token represents a session token.
type Token struct {
	Token  string `json:"token"`
	Expite string `json:"expire,omitempty"`
}
