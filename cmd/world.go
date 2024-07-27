package cmd

import (
	"github.com/rm-ryou/goPlayGround/model"
)

func CreateObjectWorld() model.HittableList {
	world := new(model.HittableList)
	world.AddObject(model.Sphere{model.Vec3D{0, 0, -1}, 0.5})
	world.AddObject(model.Sphere{model.Vec3D{0, -100.5, -1}, 100})

	return *world
}