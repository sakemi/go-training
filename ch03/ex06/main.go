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
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := ((float64(py)/height*(ymax-ymin) + ymin) + (float64(py+1)/height*(ymax-ymin) + ymin)) / 2.0
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := ((float64(px)/width*(xmax-xmin) + xmin) + (float64(px+1)/width*(xmax-xmin) + xmin)) / 2.0

			z1 := complex(x1, y1)
			z2 := complex(x2, y1)
			z3 := complex(x1, y2)
			z4 := complex(x2, y2)

			m1 := mandelbrot(z1, 0, 0, 0)
			m2 := mandelbrot(z2, 0, 0, 0)
			m3 := mandelbrot(z3, 0, 0, 0)
			m4 := mandelbrot(z4, 0, 0, 0)
			img.Set(px, py, average(m1, m2, m3, m4))
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

func average(c1, c2, c3, c4 color.Color) color.Color {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	r3, g3, b3, _ := c3.RGBA()
	r4, g4, b4, _ := c4.RGBA()

	r := uint8((r1 + r2 + r3 + r4) / 4 / 257)
	g := uint8((g1 + g2 + g3 + g4) / 4 / 257)
	b := uint8((b1 + b2 + b3 + b4) / 4 / 257)
	return color.RGBA{r, g, b, 0xff}
}
