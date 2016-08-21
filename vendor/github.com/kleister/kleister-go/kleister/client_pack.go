package kleister

// ClientPack represents a client pack API response.
type ClientPack struct {
	Client *Client `json:"client,omitempty"`
	Pack   *Pack   `json:"pack,omitempty"`
}
