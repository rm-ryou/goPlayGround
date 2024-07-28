package model

type imageEnv struct {
	aspectRatio     float64
	imageWidth      int
	viewPortHeight  float64
	samplesPerPixel int
	maxDepth        int
}

var (
	ImageEnv = imageEnv{
		aspectRatio:     16.0 / 9.0,
		imageWidth:      400,
		viewPortHeight:  2.0,
		samplesPerPixel: 100,
		maxDepth:        10,
	}
)

func (ie imageEnv) AspectRatio() float64 {
	return ie.aspectRatio
}

func (ie imageEnv) ImageWidth() int {
	return ie.imageWidth
}

func (ie imageEnv) ImageHeight() int {
	height := int(float64(ie.imageWidth) / ie.aspectRatio)

	if height < 1 {
		return 1
	} else {
		return height
	}
}

func (ie imageEnv) ViewPortWidth() float64 {
	return ie.viewPortHeight * (float64(ie.imageWidth) / float64(ie.ImageHeight()))
}

func (ie imageEnv) ViewPortHeight() float64 {
	return ie.viewPortHeight
}

func (ie imageEnv) SamplesPerPixel() int {
	return ie.samplesPerPixel
}

func (ie imageEnv) PixelSampleScale() float64 {
	return 1 / float64(ie.samplesPerPixel)
}

func (ie imageEnv) MaxDepth() int {
	return ie.maxDepth
}
