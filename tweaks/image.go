// Generating an image in golang
package main

import (
    "os"
    "fmt"
    "math"
    "image"
    "image/png"
    "image/color"
)

type Circle struct {
    X, Y, R float64
}

func (c *Circle) Brightness (x, y float64) uint8 {
    var dx, dy float64 = c.X - x, c.Y - y
    d := math.Sqrt(dx*dx + dy*dy)/c.R

    if d > 1 {
        return 0
    } else {
        return 255
    }
}

func main() {
    var width, height int = 1024, 768
    var hw, hh float64 = float64(width/2), float64(height/2)
    r := 40.0
    theta := 2*math.Pi/3
    cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
    cg := &Circle{hw - r*math.Sin(theta), hh - r*math.Cos(theta), 60}
    cb := &Circle{hw - r*math.Sin(-theta), hh - r*math.Cos(-theta), 60}

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            c := color.RGBA{
                cr.Brightness(float64(x), float64(y)),
                cg.Brightness(float64(x), float64(y)),
                cb.Brightness(float64(x), float64(y)),
                255,
            }
            img.Set(x, y, c)
        }
    }

    f, err := os.OpenFile("circles.png", os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    png.Encode(f, img)
}
