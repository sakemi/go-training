package main

import "testing"

func TestCMoF(t *testing.T) {
	if actual := MToF(1); actual != 3.28 {
		t.Error(Meters(1), actual)
	}
}

func TestFToM(t *testing.T) {
	if actual := FToM(3.28); actual != 1 {
		t.Error(Feet(3.28), actual)
	}
}
