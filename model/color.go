package model

import (
	"strconv"
	"strings"
)

type Color struct {
	Vec3D
}

func (c Color) WriteColor() string {
	var color strings.Builder

	r := int(255.999 * c.X)
	g := int(255.999 * c.Y)
	b := int(255.999 * c.Z)
	color.WriteString(strconv.Itoa(r) + " " +
		strconv.Itoa(g) + " " +
		strconv.Itoa(b) + "\n")
	return color.String()
}
