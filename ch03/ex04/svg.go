package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func writeSVG(out io.Writer, param *svgParam) {
	//svg := []string{"<svg xmlns='http://www.w3.org/2000/svg' ", "style='stroke: ", param.stroke, "; fill: ", param.fill, "; stroke-width: 0.7' width='", param.width, "' height='", param.hight, "'>"}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%s' height='%s'>", param.stroke, param.fill, param.width, param.hight)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, param)
			bx, by := corner(i, j, param)
			cx, cy := corner(i, j+1, param)
			dx, dy := corner(i+1, j+1, param)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, p *svgParam) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	width, _ := strconv.ParseFloat(p.width, 64)
	height, _ := strconv.ParseFloat(p.hight, 64)
	// width := 600.0
	// height := 320.0
	xyscale := width / 2 / xyrange
	zscale := height * 0.4

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
