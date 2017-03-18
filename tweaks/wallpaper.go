package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width, height := 1920, 1080
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{63, 63, 63, 255})
		}
	}
	f, _ := os.OpenFile("test.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
