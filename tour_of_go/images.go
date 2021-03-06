package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() (color.Model) {
	return color.RGBAModel
}

func (i Image) Bounds() (image.Rectangle) {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) (color.Color) {
	r := uint8(x * y)
	g := uint8((x + y) * 2)
	b := uint8((x - y) * 2)
	return color.RGBA{r, g, b, 0xff}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

