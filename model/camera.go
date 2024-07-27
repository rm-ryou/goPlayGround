package model

import "fmt"

type camera struct {
	focalLength float64
	center      Vec3D
}

var (
	Camera = camera{
		focalLength: 1,
		center:      Vec3D{0, 0, 0},
	}
)

func (c camera) Center() Vec3D {
	return c.center
}

func (c camera) viewportU() Vec3D {
	return Vec3D{
		X: ImageEnv.ViewPortWidth(),
		Y: 0,
		Z: 0,
	}
}

func (c camera) viewportV() Vec3D {
	return Vec3D{
		X: 0,
		Y: -ImageEnv.ViewPortHeight(),
		Z: 0,
	}
}

func (c camera) PixelDeltaU() Vec3D {
	return c.viewportU().DivNum(float64(ImageEnv.ImageWidth()))
}

func (c camera) PixelDeltaV() Vec3D {
	return c.viewportV().DivNum(float64(ImageEnv.ImageHeight()))
}

func (c camera) ViewportUpperLeft() Vec3D {
	vecCameraToViewPort := c.center.Sub(Vec3D{0, 0, c.focalLength})
	vecHalfWidth := c.viewportU().DivNum(2)
	vecHalfHeight := c.viewportV().DivNum(2)

	return vecCameraToViewPort.Sub(vecHalfWidth).Sub(vecHalfHeight)
}

func (c camera) Pixel00Loc() Vec3D {
	pixelDelta := c.PixelDeltaU().Add(c.PixelDeltaV())

	return c.ViewportUpperLeft().Add(pixelDelta.MultiNum(0.5))
}

func sampleSquare(isTest bool) Vec3D {
	fmt.Println(isTest)
	if isTest {
		return Vec3D{}
	}
	return Vec3D{RandomFloat() - 0.5, RandomFloat() - 0.5, 0}
}

func (c camera) CalcRay(x, y float64, isTest bool) Ray {
	offset := sampleSquare(isTest)
	randomPixelU := c.PixelDeltaU().MultiNum(x + offset.X)
	randomPixelV := c.PixelDeltaV().MultiNum(y + offset.Y)
	randomPixelUV := randomPixelU.Add(randomPixelV)
	pixelSample := c.Pixel00Loc().Add(randomPixelUV)
	rayOrig := c.center
	rayDir := pixelSample.Sub(rayOrig)

	return Ray{rayOrig, rayDir}
}
