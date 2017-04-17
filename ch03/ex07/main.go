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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return rootColor(z, 255-contrast*i)
		}
	}
	return color.Black
}

func rootColor(z complex128, c uint8) color.Color {
	r := real(z)
	i := imag(z)
	if r > -1.0 {
		if i > 0 {
			return color.RGBA{255 - c, 0, 0, 255}
		}
		return color.RGBA{0, 255 - c, 0, 255}
	}
	if i > 0 {
		return color.RGBA{0, 0, 255 - c, 255}
	}
	return color.Gray{255 - c}
}
