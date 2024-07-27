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
	intensity := Interval{0.999, 0}
	r := int(256 * intensity.Clamp(c.R))
	g := int(256 * intensity.Clamp(c.G))
	b := int(256 * intensity.Clamp(c.B))

	color.WriteString(strconv.Itoa(r) + " " +
		strconv.Itoa(g) + " " +
		strconv.Itoa(b) + "\n")

	return color.String()
}
