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

// TODO: use interface
func (c Color) Add(other Color) Color {
	return Color{
		R: c.R + other.R,
		G: c.G + other.G,
		B: c.B + other.B,
	}
}

func (c Color) Multi(other Color) Color {
	return Color{
		R: c.R * other.R,
		G: c.G * other.G,
		B: c.B * other.B,
	}
}

func (c Color) MultiNum(n float64) Color {
	return Color{
		R: c.R * n,
		G: c.G * n,
		B: c.B * n,
	}
}
