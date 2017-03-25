package tempconv

import "testing"

func TestCToK(t *testing.T) {
	if actual := CToK(0); actual != 273.15 {
		t.Error(Celsius(0), actual)
	}

	if actual := CToK(AbsoluteZeroC); actual != 0 {
		t.Error(AbsoluteZeroC, actual)
	}

	if actual := CToK(-40); actual != 233.15 {
		t.Error(Celsius(-40), actual)
	}
}

func TestKToC(t *testing.T) {
	if actual := KToC(0); actual != -273.15 {
		t.Error(Kelvin(0), actual)
	}

	if actual := KToC(273.15); actual != 0 {
		t.Error(Kelvin(273.15), actual)
	}

	if actual := CToK(-40); actual != 233.15 {
		t.Error(Celsius(-40), actual)
	}
}

func TestFToK(t *testing.T) {
	if actual := FToK(32); actual != 273.15 {
		t.Error(Fahrenheit(32), actual)
	}

	if actual := FToK(-40); actual != 233.15 {
		t.Error(Fahrenheit(-40), actual)
	}
}

func TestKToF(t *testing.T) {
	if actual := KToF(273.15); actual != 32 {
		t.Error(Kelvin(273.15), actual)
	}

	if actual := KToF(233.15); actual != -40 {
		t.Error(Kelvin(233.15), actual)
	}
}

func TestFToC(t *testing.T) {
	if actual := FToC(-40); actual != -40 {
		t.Error(Fahrenheit(-40), actual)
	}

	if actual := FToC(32); actual != 0 {
		t.Error(Fahrenheit(32), actual)
	}
}

func TestCToF(t *testing.T) {
	if actual := CToF(0); actual != 32 {
		t.Error(Celsius(0), actual)
	}

	if actual := CToF(-40); actual != -40 {
		t.Error(Celsius(-40), actual)
	}
}

func TestRound(t *testing.T) {
	if actual := round(233.14999999999998, .5, 2); actual != 233.15 {
		t.Error(actual)
	}
}
