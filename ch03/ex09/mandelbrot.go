package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"strconv"
)

const (
	xminDef, yminDef, xmaxDef, ymaxDef = -2, -2, +2, +2
	width, height                      = 1024, 1024
	c64, c128, bf, rat                 = "64", "128", "bf", "rat"
)

func renderMandelbrot(out io.Writer, param *fractalParam) {
	scale, err := strconv.ParseFloat(param.scale, 64)
	if err != nil {
		scale = 1.0
	}
	if scale <= 0 {
		//if illegal argument is given, do nothing
		return
	}
	x, err := strconv.ParseFloat(param.x, 64)
	if err != nil {
		return
	}
	y, err := strconv.ParseFloat(param.y, 64)
	if err != nil {
		return
	}
	xmin := xminDef/scale + x
	ymin := yminDef/scale + y
	xmax := xmaxDef/scale + x
	ymax := ymaxDef/scale + y

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
