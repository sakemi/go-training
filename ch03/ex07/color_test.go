package main

import (
	"image/color"
	"testing"
)

func TestRootColor(t *testing.T) {
	z := complex(-0.9, 1)
	expected := color.RGBA{255, 0, 0, 255}
	if actual := rootColor(z, 0); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z = complex(-0.9, -1)
	expected = color.RGBA{0, 255, 0, 255}
	if actual := rootColor(z, 0); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z = complex(-1.1, 1)
	expected = color.RGBA{0, 0, 255, 255}
	if actual := rootColor(z, 0); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z2 := complex(-1.1, -1)
	expected2 := color.Gray{255}
	if actual := rootColor(z2, 0); actual != expected2 {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}

func TestRootColorContrast(t *testing.T) {
	z := complex(-0.9, 1)
	expected := color.RGBA{245, 0, 0, 255}
	if actual := rootColor(z, 10); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z = complex(-0.9, -1)
	expected = color.RGBA{0, 245, 0, 255}
	if actual := rootColor(z, 10); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z = complex(-1.1, 1)
	expected = color.RGBA{0, 0, 245, 255}
	if actual := rootColor(z, 10); actual != expected {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}

	z2 := complex(-1.1, -1)
	expected2 := color.Gray{245}
	if actual := rootColor(z2, 10); actual != expected2 {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
