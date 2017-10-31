package sdk

import (
	"github.com/hashicorp/go-version"
)

// MinecraftBySlug sorts a list of minecraft by slug field.
type MinecraftBySlug []*Minecraft

// Len is part of the minecraft sorting algorithm.
func (u MinecraftBySlug) Len() int {
	return len(u)
}

// Swap is part of the minecraft sorting algorithm.
func (u MinecraftBySlug) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the minecraft sorting algorithm.
func (u MinecraftBySlug) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Slug)
	v2, err2 := version.NewVersion(u[j].Slug)

	if err1 != nil || err2 != nil {
		return u[i].Slug < u[j].Slug
	}

	return v1.LessThan(v2)
}
