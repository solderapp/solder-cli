package kleister

// BuildVersion represents a build version API response.
type BuildVersion struct {
	Build   *Build   `json:"build,omitempty"`
	Version *Version `json:"version,omitempty"`
}
