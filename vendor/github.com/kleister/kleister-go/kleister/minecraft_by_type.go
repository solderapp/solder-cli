package kleister

import (
	"github.com/hashicorp/go-version"
)

// MinecraftByType sorts a list of minecraft by type field.
type MinecraftByType []*Minecraft

// Len is part of the minecraft sorting algorithm.
func (u MinecraftByType) Len() int {
	return len(u)
}

// Swap is part of the minecraft sorting algorithm.
func (u MinecraftByType) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the minecraft sorting algorithm.
func (u MinecraftByType) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Type)
	v2, err2 := version.NewVersion(u[j].Type)

	if err1 != nil || err2 != nil {
		return u[i].Type < u[j].Type
	}

	return v1.LessThan(v2)
}
