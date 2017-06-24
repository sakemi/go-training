package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {
	var routine int
	if len(os.Args) == 1 {
		routine = 1
	} else {
		r, err := strconv.Atoi(os.Args[1])
		if err != nil || r < 1 {
			routine = 1
		} else if r > 1024 {
			routine = 1024
		} else {
			routine = r
		}
	}
	img := calc(routine)
	png.Encode(os.Stdout, img)
}

type pix struct {
	px  int
	py  int
	col color.Color
}

func calc(routine int) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	ch := make(chan []pix, routine)
	rang := height / routine
	var start, end int
	for i := 0; i < routine; i++ {
		start = rang * i
		if i != routine-1 {
			end = start + rang
		} else {
			end = height
		}

		go func(start, end int) {
			result := []pix{}
			for py := start; py < end; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					result = append(result, pix{px, py, mandelbrot(z)})
				}
			}
			ch <- result
		}(start, end)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < routine; i++ {
		p := <-ch
		for _, v := range p {
			img.Set(v.px, v.py, v.col)
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
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

// func acos(z complex128) color.Color {
// 	v := cmplx.Acos(z)
// 	blue := uint8(real(v)*128) + 127
// 	red := uint8(imag(v)*128) + 127
// 	return color.YCbCr{192, blue, red}
// }
//
// func sqrt(z complex128) color.Color {
// 	v := cmplx.Sqrt(z)
// 	blue := uint8(real(v)*128) + 127
// 	red := uint8(imag(v)*128) + 127
// 	return color.YCbCr{128, blue, red}
// }
//
// func newton(z complex128) color.Color {
// 	const iterations = 37
// 	const contrast = 7
// 	for i := uint8(0); i < iterations; i++ {
// 		z -= (z - 1/(z*z*z)) / 4
// 		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
// 			return color.Gray{255 - contrast*i}
// 		}
// 	}
// 	return color.Black
// }
