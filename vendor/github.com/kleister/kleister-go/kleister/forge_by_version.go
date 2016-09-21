package kleister

import (
	"github.com/hashicorp/go-version"
)

// ForgeByVersion sorts a list of forge by version field.
type ForgeByVersion []*Forge

// Len is part of the forge sorting algorithm.
func (u ForgeByVersion) Len() int {
	return len(u)
}

// Swap is part of the forge sorting algorithm.
func (u ForgeByVersion) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the forge sorting algorithm.
func (u ForgeByVersion) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Version)
	v2, err2 := version.NewVersion(u[j].Version)

	if err1 != nil || err2 != nil {
		return u[i].Version < u[j].Version
	}

	return v1.LessThan(v2)
}
