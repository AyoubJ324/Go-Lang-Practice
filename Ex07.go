package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x1, y1 int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x1, i.y1)
}

func (i Image) At(x, y int) color.Color {
	output := uint8(x * y)
	return color.RGBA{output, output, 255, 255}
}

func main() {
	m := Image{256, 65}
	pic.ShowImage(m)
}
