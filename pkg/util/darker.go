package util

import (
	"image"
	"image/color"
)

func Darker(img image.Image, percent float64) image.Image {
	bounds := img.Bounds()
	darkened := image.NewRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			// Reduce each color component by the specified percentage
			r -= uint32(float64(r) * percent / 100)
			g -= uint32(float64(g) * percent / 100)
			b -= uint32(float64(b) * percent / 100)
			darkened.Set(x, y, color.RGBA{uint8(r / 256), uint8(g / 256), uint8(b / 256), uint8(a / 256)})
		}
	}
	return darkened
}
