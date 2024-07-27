package model

import "math"

type Vec3D struct {
	X, Y, Z float64
}

func (v Vec3D) Add(other Vec3D) Vec3D {
	return Vec3D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3D) Sub(other Vec3D) Vec3D {
	return Vec3D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3D) MultiVec(other Vec3D) Vec3D {
	return Vec3D{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

func (v Vec3D) DivVec(other Vec3D) Vec3D {
	return Vec3D{
		X: v.X / other.X,
		Y: v.Y / other.Y,
		Z: v.Z / other.Z,
	}
}

func (v Vec3D) MultiNum(num float64) Vec3D {
	return Vec3D{
		X: v.X * num,
		Y: v.Y * num,
		Z: v.Z * num,
	}
}

func (v Vec3D) DivNum(num float64) Vec3D {
	return Vec3D{
		X: v.X / num,
		Y: v.Y / num,
		Z: v.Z / num,
	}
}

func (v Vec3D) Dot(other Vec3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3D) Cross(other Vec3D) Vec3D {
	return Vec3D{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3D) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3D) Norm() Vec3D {
	norm := v.Length()
	return v.DivNum(norm)
}
