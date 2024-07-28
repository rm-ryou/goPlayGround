package model

type Hittable interface {
	Hit(ray Ray, rayTMax, rayTMin float64, rec HitRecord) bool
}
