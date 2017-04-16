package main

import (
	"image/color"
	"testing"
)

func TestMandelbrotBlack(t *testing.T) {
	z := complex(0, 0)

	if actual := mandelbrot(z, 0, 0, 0); actual != color.Black {
		t.Errorf("z:%d+%di return:%v", 0, 0, actual)
	}
}

func TestMandelbrotColored(t *testing.T) {
	z := complex(1, 0)
	expected := color.RGBA{127, 127, 127, 255}
	if actual := mandelbrot(z, 0, 0, 0); actual != expected {
		t.Errorf("z:%d+%di return:%v", 2, 2, actual)
	}
}

func TestMandelbrotNoised(t *testing.T) {
	z := complex(1, 0)
	expected := color.RGBA{137, 147, 157, 255}
	if actual := mandelbrot(z, 10, 20, 30); actual != expected {
		t.Errorf("z:%d+%di return:%v", 2, 2, actual)
	}
}
