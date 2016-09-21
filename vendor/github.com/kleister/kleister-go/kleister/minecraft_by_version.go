package kleister

import (
	"github.com/hashicorp/go-version"
)

// MinecraftByVersion sorts a list of minecraft by version field.
type MinecraftByVersion []*Minecraft

// Len is part of the minecraft sorting algorithm.
func (u MinecraftByVersion) Len() int {
	return len(u)
}

// Swap is part of the minecraft sorting algorithm.
func (u MinecraftByVersion) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the minecraft sorting algorithm.
func (u MinecraftByVersion) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Version)
	v2, err2 := version.NewVersion(u[j].Version)

	if err1 != nil || err2 != nil {
		return u[i].Version < u[j].Version
	}

	return v1.LessThan(v2)
}
