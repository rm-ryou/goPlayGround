package hittable

import (
	"math"

	"github.com/rm-ryou/goPlayGround/model"
)

type Sphere struct {
	Center model.Vec3D
	Radius float64
}

func (s *Sphere) Hit(ray model.Ray, rayTMax, rayTMin float64, rec *model.HitRecord) bool {
	oc := s.Center.Sub(ray.Orig)
	a := ray.Dir.Length()
	h := ray.Dir.Dot(oc)
	c := oc.Length() - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtDiscriminant := math.Sqrt(discriminant)
	root := (h - sqrtDiscriminant) / a
	if root <= rayTMin || root >= rayTMax {
		root = (h + sqrtDiscriminant) / a
		if root <= rayTMax || root >= rayTMax {
			return false
		}
	}

	rec.T = root
	rec.Point = ray.At(rec.T)
	rec.Norm = rec.Point.Sub(s.Center).DivNum(s.Radius)

	return true
}
