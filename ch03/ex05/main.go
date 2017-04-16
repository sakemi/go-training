package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z, 0, 20, 100))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128, r, g, b uint8) color.Color {
	const (
		iterations = 200
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			strength := uint8(255 * (1.0 - 1.0/float64(n)))
			return color.RGBA{R: (strength + r) % 255, G: (strength + g) % 255, B: (strength + b) % 255, A: 0xff}
		}
	}
	return color.Black
}
