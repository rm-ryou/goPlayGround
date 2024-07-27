package model

type Hittable interface {
	Hit(ray Ray, rayT Interval, rec HitRecord) bool
}
