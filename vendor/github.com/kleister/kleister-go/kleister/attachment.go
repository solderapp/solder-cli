package kleister

import (
	"time"
)

// Attachment represents a attachment API response.
type Attachment struct {
	URL       string    `json:"url,omitempty"`
	MD5       string    `json:"md5,omitempty"`
	Upload    string    `json:"upload,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Attachment) String() string {
	return s.URL
}
