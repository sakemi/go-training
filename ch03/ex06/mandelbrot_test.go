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

func TestAverage(t *testing.T) {
	c := color.RGBA{0, 0, 0, 255}
	if actual := average(c, c, c, c); actual != c {
		t.Errorf("input:%v average:%v", c, actual)
	}

	c1 := color.RGBA{0, 0, 0, 255}
	c2 := color.RGBA{10, 20, 30, 255}
	c3 := color.RGBA{100, 30, 50, 255}
	c4 := color.RGBA{50, 30, 122, 255}
	expected := color.RGBA{40, 20, 50, 255}
	if actual := average(c1, c2, c3, c4); actual != expected {
		t.Errorf("input:%v,%v,%v,%v average:%v", c1, c2, c3, c4, actual)
	}
}
