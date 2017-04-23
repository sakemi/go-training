package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		c64, c128, bf, rat     = "64", "128", "bf", "rat"
	)

	args := os.Args[1:]
	mode := args[0]

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	switch mode {
	case c64:
		for py := 0; py < height; py++ {
			y := float32(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float32(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot64(z))
			}
		}
	case c128:
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot128(z))
			}
		}
	case bf:
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				img.Set(px, py, mandelbrotBigFloat128(x, y))
			}
		}
	case rat:
		for py := 0; py < height; py++ {
			y := big.NewRat(int64(py)*(ymax-ymin), height)
			y.Add(y, big.NewRat(ymin, 1))
			for px := 0; px < width; px++ {
				x := big.NewRat(int64(px)*(xmax-xmin), width)
				x.Add(x, big.NewRat(xmin, 1))
				img.Set(px, py, mandelbrotRat(x, y))
			}
		}
	}
	png.Encode(os.Stdout, img)
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

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		dr := real(v) * real(v)
		di := imag(v) * imag(v)
		if math.Sqrt(float64(dr+di)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

const prec = 128

func mandelbrotBigFloat128(x, y float64) color.Color {
	const iterations = 100
	const contrast = 15

	two := new(big.Float).SetPrec(prec).SetInt64(2)
	bx := new(big.Float).SetPrec(prec).SetFloat64(x)
	by := new(big.Float).SetPrec(prec).SetFloat64(y)

	vr := new(big.Float).SetPrec(prec).SetFloat64(0)
	vi := new(big.Float).SetPrec(prec).SetFloat64(0)
	for n := uint8(0); n < iterations; n++ {
		vr, vi = floatFunc(vr, vi, bx, by)
		vr2 := new(big.Float).Mul(vr, vr)
		vi2 := new(big.Float).Mul(vi, vi)
		squareSum := new(big.Float).Add(vr2, vi2)
		abs := sqrt(squareSum)
		if abs.Cmp(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(x, y *big.Rat) color.Color {
	const iterations = 15
	const contrast = 25

	two := big.NewRat(2, 1)

	vr := new(big.Rat)
	vi := new(big.Rat)
	for n := uint8(0); n < iterations; n++ {
		vr, vi = ratFunc(vr, vi, x, y)
		vr2 := new(big.Rat).Mul(vr, vr)
		vi2 := new(big.Rat).Mul(vi, vi)
		squareSum := new(big.Rat).Add(vr2, vi2)
		abs := sqrtRat(squareSum)
		if abs.Cmp(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func multiBigComplex(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	vr2 := new(big.Float).Mul(xr, yr)
	vi2 := new(big.Float).Mul(xi, yi)
	ri := new(big.Float).Mul(xr, yi)
	ir := new(big.Float).Mul(xi, yr)

	r := new(big.Float)
	r.Sub(vr2, vi2)

	i := new(big.Float)
	i.Add(ri, ir)

	return r, i
}

func sumBigComplex(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	r := new(big.Float)
	r.Add(xr, yr)
	i := new(big.Float)
	i.Add(xi, yi)

	return r, i
}

func floatFunc(vr, vi, zr, zi *big.Float) (*big.Float, *big.Float) {
	v2r, v2i := multiBigComplex(vr, vi, vr, vi)
	r, i := sumBigComplex(v2r, v2i, zr, zi)

	return r, i
}

func sqrt(x *big.Float) *big.Float {
	steps := int(math.Log2(prec))
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	tmp := new(big.Float)
	result := new(big.Float).SetPrec(prec).SetInt64(1)
	for i := 0; i <= steps; i++ {
		tmp.Quo(x, result)
		tmp.Add(result, tmp)
		result.Mul(half, tmp)
	}
	return result
}

func ratFunc(vr, vi, zr, zi *big.Rat) (*big.Rat, *big.Rat) {
	v2r, v2i := multiRatComplex(vr, vi, vr, vi)
	r, i := sumRatComplex(v2r, v2i, zr, zi)

	return r, i
}

func sumRatComplex(xr, xi, yr, yi *big.Rat) (*big.Rat, *big.Rat) {
	r := new(big.Rat)
	r.Add(xr, yr)
	i := new(big.Rat)
	i.Add(xi, yi)

	return r, i
}

func multiRatComplex(xr, xi, yr, yi *big.Rat) (*big.Rat, *big.Rat) {
	vr2 := new(big.Rat).Mul(xr, yr)
	vi2 := new(big.Rat).Mul(xi, yi)
	ri := new(big.Rat).Mul(xr, yi)
	ir := new(big.Rat).Mul(xi, yr)

	r := new(big.Rat)
	r.Sub(vr2, vi2)

	i := new(big.Rat)
	i.Add(ri, ir)

	return r, i
}

func sqrtRat(x *big.Rat) *big.Rat {
	steps := int(math.Log2(prec))
	half := big.NewRat(1, 2)
	tmp := new(big.Rat)
	result := big.NewRat(1, 1)
	for i := 0; i <= steps; i++ {
		tmp.Quo(x, result)
		tmp.Add(result, tmp)
		result.Mul(half, tmp)
	}
	return result
}
