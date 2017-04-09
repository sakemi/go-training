package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6

	peak   = 1
	valley = -1
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, i1 := corner(i+1, j)
			bx, by, i2 := corner(i, j)
			cx, cy, i3 := corner(i, j+1)
			dx, dy, i4 := corner(i+1, j+1)
			if i1 == peak || i2 == peak || i3 == peak || i4 == peak {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#ff0000'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else if i1 == valley || i2 == valley || i3 == valley || i4 == valley {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#0000ff'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, int) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	inf := checkInfrection(z, i, j)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, inf
}

func checkInfrection(z float64, i, j int) int {
	if i == 0 || j == 0 {
		return 0
	}

	x := xyrange * (float64(i-1)/cells - 0.5)
	y := xyrange * (float64(j-1)/cells - 0.5)
	prev := f(x, y)

	x = xyrange * (float64(i+1)/cells - 0.5)
	y = xyrange * (float64(j+1)/cells - 0.5)
	next := f(x, y)

	if z > prev && z > next {
		return peak
	}

	if z < prev && z < next {
		return valley
	}

	return 0
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
