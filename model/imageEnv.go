package model

type imageEnv struct {
	aspectRatio float64
	imageWidth int
	viewPortHeight float64
}

var (ImageEnv = imageEnv{
	aspectRatio: 16 / 9,
	imageWidth: 400,
	viewPortHeight: 2,
})

func (ie imageEnv)AspectRatio() float64 {
	return ie.aspectRatio
}

func (ie imageEnv)ImageWidth() int {
	return ie.imageWidth
}

func (ie imageEnv)ImageHeight() int {
	height := int(float64(ie.imageWidth / int(ie.aspectRatio)))

	if height < 1 {
		return 1
	} else {
		return height
	}
}

func (ie imageEnv)ViewPortWidth() float64 {
	return ie.viewPortHeight * (float64(ie.imageWidth) / float64(ie.ImageHeight()))
}

func (ie imageEnv)ViewPortHeight() float64 {
	return ie.viewPortHeight
}
