package main

import (
	"math"
	"testing"
)

func TestF(t *testing.T) {
	if actual := f(0, 0); actual != 1 {
		t.Error(actual)
	}

	if actual := f(-20, -10); actual != -1 {
		t.Error(actual)
	}
}

func TestFNaN(t *testing.T) {
	if actual := f(-15, -15); !math.IsNaN(actual) {
		t.Error(actual)
	}
}

func TestCorner(t *testing.T) {
	_, _, ok := corner(0, 0)
	if ok {
		t.Error("ok is true even though f() returns NaN")
	}
}
