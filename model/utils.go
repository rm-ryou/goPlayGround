package model

import (
	"math"
	"math/rand"
)

func DegreeToRad(degree float64) float64 {
	return degree * math.Pi / 180
}

func RandomFloat() float64 {
	return rand.Float64()
}

func RandomFloatWithIn(max, min float64) float64 {
	return min + (max-min)*RandomFloat()
}
