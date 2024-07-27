package hittable

import "github.com/rm-ryou/goPlayGround/model"

type Hittable interface {
	Hit(ray model.Ray, rayTMax, rayTMin float64, rec model.HitRecord) bool
}
