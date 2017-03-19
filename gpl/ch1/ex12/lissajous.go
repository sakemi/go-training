package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, param *lissajousParam) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: param.nframes}
	phase := 0.0
	for i := 0; i < param.nframes; i++ {
		rect := image.Rect(0, 0, 2*param.size+1, 2*param.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < param.cycles*2*math.Pi; t += param.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(param.size+int(x*float64(param.size)+0.5), param.size+int(y*float64(param.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, param.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
