package sdk

import (
	"github.com/hashicorp/go-version"
)

// ForgeBySlug sorts a list of forge by slug field.
type ForgeBySlug []*Forge

// Len is part of the forge sorting algorithm.
func (u ForgeBySlug) Len() int {
	return len(u)
}

// Swap is part of the forge sorting algorithm.
func (u ForgeBySlug) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the forge sorting algorithm.
func (u ForgeBySlug) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Slug)
	v2, err2 := version.NewVersion(u[j].Slug)

	if err1 != nil || err2 != nil {
		return u[i].Slug < u[j].Slug
	}

	return v1.LessThan(v2)
}
