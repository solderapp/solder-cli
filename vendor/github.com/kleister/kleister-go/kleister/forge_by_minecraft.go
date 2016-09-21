package kleister

import (
	"github.com/hashicorp/go-version"
)

// ForgeByMinecraft sorts a list of forge by minecraft field.
type ForgeByMinecraft []*Forge

// Len is part of the forge sorting algorithm.
func (u ForgeByMinecraft) Len() int {
	return len(u)
}

// Swap is part of the forge sorting algorithm.
func (u ForgeByMinecraft) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less is part of the forge sorting algorithm.
func (u ForgeByMinecraft) Less(i, j int) bool {
	v1, err1 := version.NewVersion(u[i].Minecraft)
	v2, err2 := version.NewVersion(u[j].Minecraft)

	if err1 != nil || err2 != nil {
		return u[i].Minecraft < u[j].Minecraft
	}

	return v1.LessThan(v2)
}
