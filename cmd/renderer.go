package cmd

import (
	"strconv"
	"strings"

	"github.com/rm-ryou/goPlayGround/model"
)

func createFormat() string {
	var format strings.Builder
	width := model.ImageEnv.ImageWidth()
	height := model.ImageEnv.ImageHeight()

	format.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	return format.String()
}

func createImage() string {
	var image strings.Builder
	width := model.ImageEnv.ImageWidth()
	height := model.ImageEnv.ImageHeight()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			color := model.Color{
				R: float64(j) / (float64(width) - 1),
				G: float64(i) / (float64(height) - 1),
				B: 0,
			}

			image.WriteString(color.WriteColor())
		}
	}
	return image.String()
}

func CreateData() string {
	return createFormat() + createImage()
}
