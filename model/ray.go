package model

import (
	"math"
)

type Ray struct {
	Orig Vec3D
	Dir  Vec3D
}

func (r Ray) At(t float64) Vec3D {
	return r.Orig.Add(r.Dir.MultiNum(t))
}

func (r Ray) Color(worldObjects HittableList, isTest bool) Color {
	hitRec := new(HitRecord)
	if worldObjects.Hit(r, Interval{math.Inf(1), 0}, hitRec) {
		dir := hitRec.Norm.GenRandomHemiSphere(isTest)
		if isTest { return dir.MultiNum(0.5).Vec3DToColor() }
		return Ray{hitRec.Point, dir}.Color(worldObjects, isTest).MultiNum(0.5)
	}

	unitDir := r.Dir.Norm()
	a := 0.5 * (unitDir.Y + 1)
	vec := Vec3D{1, 1, 1}.MultiNum(1 - a).Add(Vec3D{0.5, 0.7, 1}.MultiNum(a))

	return vec.Vec3DToColor()
}
