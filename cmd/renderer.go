package cmd

import (
	"strconv"
	"strings"

	"github.com/rm-ryou/goPlayGround/model"
)

func createFormat(width, height int) string {
	var format strings.Builder

	format.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	return format.String()
}

func createImage(width, height int) string {
	var image strings.Builder

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pixDeltaU := model.Camera.PixelDeltaU().MultiNum(float64(j))
			pixDeltaV := model.Camera.PixelDeltaV().MultiNum(float64(i))
			pixCenter := model.Camera.Pixel00Loc().Add(pixDeltaU.Add(pixDeltaV))
			rayDir := pixCenter.Sub(model.Camera.Center())

			ray := model.Ray{model.Camera.Center(), rayDir}

			color := ray.Color()
			image.WriteString(color.WriteColor())
		}
	}
	return image.String()
}

func CreateData() string {
	width := model.ImageEnv.ImageWidth()
	height := model.ImageEnv.ImageHeight()

	return createFormat(width, height) + createImage(width, height)
}
