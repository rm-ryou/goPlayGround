package model

import (
	"strconv"
	"strings"
)

type Color struct {
	R, G, B float64
}

func (c Color) WriteColor() string {
	var color strings.Builder

	r := int(255.999 * c.R)
	g := int(255.999 * c.G)
	b := int(255.999 * c.B)
	color.WriteString(strconv.Itoa(r) + " " +
		strconv.Itoa(g) + " " +
		strconv.Itoa(b) + "\n")
	return color.String()
}
