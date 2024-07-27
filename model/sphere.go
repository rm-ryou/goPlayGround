package model

import (
	"math"
)

type Sphere struct {
	Center Vec3D
	Radius float64
}

func (s *Sphere) Hit(ray Ray, rayTMax, rayTMin float64, rec *HitRecord) bool {
	oc := s.Center.Sub(ray.Orig)
	a := ray.Dir.Dot(ray.Dir)
	h := ray.Dir.Dot(oc)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtDiscriminant := math.Sqrt(discriminant)
	root := (h - sqrtDiscriminant) / a
	if root <= rayTMin || root >= rayTMax {
		root = (h + sqrtDiscriminant) / a
		if root <= rayTMin || root >= rayTMax {
			return false
		}
	}

	rec.T = root
	rec.Point = ray.At(rec.T)
	outwardNorm := rec.Point.Sub(s.Center).DivNum(s.Radius)
	rec.SetFaceNorm(ray, outwardNorm)

	return true
}
