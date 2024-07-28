package model

import (
	"math"
	"strconv"
	"strings"
)

type Color struct {
	R, G, B float64
}

func (c Color) WriteColor() string {
	var color strings.Builder
	var colorBase float64 = 256
	intensity := Interval{0.999, 0}

	r := linearToGamma(c.R)
	g := linearToGamma(c.G)
	b := linearToGamma(c.B)

	rGammaSpace := int(colorBase * intensity.Clamp(r))
	gGammaSpace := int(colorBase * intensity.Clamp(g))
	bGammaSpace := int(colorBase * intensity.Clamp(b))

	color.WriteString(strconv.Itoa(rGammaSpace) + " " +
		strconv.Itoa(gGammaSpace) + " " +
		strconv.Itoa(bGammaSpace) + "\n")

	return color.String()
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}
	return 0
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
