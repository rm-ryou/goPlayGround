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

func createImage(width, height int, objectList model.HittableList) string {
	var image strings.Builder

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pixelColor := model.Color{}
			for sample := 0; sample < model.ImageEnv.SamplesPerPixel(); sample++ {
				ray := model.Camera.CalcRay(float64(j), float64(i), false)
				pixelColor = pixelColor.Add(ray.Color(objectList))
			}
			color := pixelColor.MultiNum(model.ImageEnv.PixelSampleScale())
			image.WriteString(color.WriteColor())
		}
	}
	return image.String()
}

func CreateData(objectList model.HittableList) string {
	width := model.ImageEnv.ImageWidth()
	height := model.ImageEnv.ImageHeight()

	return createFormat(width, height) + createImage(width, height, objectList)
}
