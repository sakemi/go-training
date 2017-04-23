package main

import (
	"math/big"
	"testing"
)

func TestSqrt(t *testing.T) {
	//setup
	x := new(big.Float).SetInt64(2)
	//execute
	actual := sqrt(x)
	//test
	diff := new(big.Float)
	diff.Mul(actual, actual)
	diff.Sub(x, diff)
	th := new(big.Float).SetFloat64(0.0001)
	if diff.Cmp(th) > 0 {
		t.Errorf("Sqrt(%v): %v", x, actual)
	}

	//setup
	x = new(big.Float).SetFloat64(3.58)
	//execute
	actual = sqrt(x)
	//test
	diff = new(big.Float)
	diff.Mul(actual, actual)
	diff.Sub(x, diff)
	if diff.Cmp(th) > 0 {
		t.Errorf("Sqrt(%v): %v", x, actual)
	}
}

func TestSumBigComplex(t *testing.T) {
	//setup
	xr := new(big.Float).SetInt64(1)
	xi := new(big.Float).SetInt64(2)
	yr := new(big.Float).SetInt64(3)
	yi := new(big.Float).SetInt64(4)
	expectR, expectI := new(big.Float).SetInt64(4), new(big.Float).SetInt64(6)

	//execute
	actualR, actualI := sumBigComplex(xr, xi, yr, yi)

	//test
	if actualR.Cmp(expectR) != 0 || actualI.Cmp(expectI) != 0 {
		t.Errorf("%v+%vi + %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}

	//setup
	xr = new(big.Float).SetFloat64(1.9)
	xi = new(big.Float).SetFloat64(2.8)
	yr = new(big.Float).SetFloat64(3.7)
	yi = new(big.Float).SetFloat64(4.6)
	expectR, expectI = new(big.Float).SetFloat64(5.6), new(big.Float).SetFloat64(7.4)

	//execute
	actualR, actualI = sumBigComplex(xr, xi, yr, yi)
	diffR := new(big.Float).Sub(actualR, expectR)
	diffI := new(big.Float).Sub(actualI, expectI)
	th := new(big.Float).SetFloat64(0.0001)

	//test
	if diffR.Cmp(th) > 0 || diffI.Cmp(th) > 0 {
		t.Errorf("%v+%vi + %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}
}

func TestMultiBigComplex(t *testing.T) {
	//setup
	xr := new(big.Float).SetInt64(1)
	xi := new(big.Float).SetInt64(2)
	yr := new(big.Float).SetInt64(3)
	yi := new(big.Float).SetInt64(4)
	expectR, expectI := new(big.Float).SetInt64(-5), new(big.Float).SetInt64(10)

	//execute
	actualR, actualI := multiBigComplex(xr, xi, yr, yi)

	//test
	if actualR.Cmp(expectR) != 0 || actualI.Cmp(expectI) != 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}

	//setup
	xr = new(big.Float).SetFloat64(5.5)
	xi = new(big.Float).SetFloat64(2.8)
	yr = new(big.Float).SetFloat64(3.7)
	yi = new(big.Float).SetFloat64(4.6)
	expectR, expectI = new(big.Float).SetFloat64(8.07), new(big.Float).SetFloat64(35.66)

	//execute
	actualR, actualI = multiBigComplex(xr, xi, yr, yi)
	diffR := new(big.Float).Sub(actualR, expectR)
	diffI := new(big.Float).Sub(actualI, expectI)
	th := new(big.Float).SetFloat64(0.0001)

	//test
	if diffR.Cmp(th) > 0 || diffI.Cmp(th) > 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}
}

func TestSumRatComplex(t *testing.T) {
	//setup
	xr := big.NewRat(1, 1)
	xi := big.NewRat(2, 1)
	yr := big.NewRat(3, 1)
	yi := big.NewRat(4, 1)
	expectR, expectI := big.NewRat(4, 1), big.NewRat(6, 1)

	//execute
	actualR, actualI := sumRatComplex(xr, xi, yr, yi)

	//test
	if actualR.Cmp(expectR) != 0 || actualI.Cmp(expectI) != 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}

	//setup
	xr = big.NewRat(1, 2)
	xi = big.NewRat(2, 3)
	yr = big.NewRat(4, 3)
	yi = big.NewRat(5, 2)
	expectR, expectI = big.NewRat(11, 6), big.NewRat(19, 6)

	//execute
	actualR, actualI = sumRatComplex(xr, xi, yr, yi)
	diffR := new(big.Rat).Sub(actualR, expectR)
	diffI := new(big.Rat).Sub(actualI, expectI)
	th := big.NewRat(1, 10000)

	//test
	if diffR.Cmp(th) > 0 || diffI.Cmp(th) > 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}
}

func TestMultiRatComplex(t *testing.T) {
	//setup
	xr := big.NewRat(1, 1)
	xi := big.NewRat(2, 1)
	yr := big.NewRat(3, 1)
	yi := big.NewRat(4, 1)
	expectR, expectI := big.NewRat(-5, 1), big.NewRat(10, 1)

	//execute
	actualR, actualI := multiRatComplex(xr, xi, yr, yi)

	//test
	if actualR.Cmp(expectR) != 0 || actualI.Cmp(expectI) != 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}

	//setup
	xr = big.NewRat(1, 2)
	xi = big.NewRat(2, 3)
	yr = big.NewRat(4, 3)
	yi = big.NewRat(5, 2)
	expectR, expectI = big.NewRat(-1, 1), big.NewRat(77, 36)

	//execute
	actualR, actualI = multiRatComplex(xr, xi, yr, yi)
	diffR := new(big.Rat).Sub(actualR, expectR)
	diffI := new(big.Rat).Sub(actualI, expectI)
	th := big.NewRat(1, 10000)

	//test
	if diffR.Cmp(th) > 0 || diffI.Cmp(th) > 0 {
		t.Errorf("%v+%vi * %v+%vi = %v+%vi", xr, xi, yr, yi, actualR, actualI)
	}
}

func TestSqrtRat(t *testing.T) {
	//setup
	x := big.NewRat(2, 1)
	//execute
	actual := sqrtRat(x)
	//test
	diff := new(big.Rat)
	diff.Mul(actual, actual)
	diff.Sub(x, diff)
	th := big.NewRat(1, 10000)
	if diff.Cmp(th) > 0 {
		t.Errorf("Sqrt(%v): %v", x, actual)
	}

	//setup
	x = big.NewRat(5, 3)
	//execute
	actual = sqrtRat(x)
	//test
	diff = new(big.Rat)
	diff.Mul(actual, actual)
	diff.Sub(x, diff)
	if diff.Cmp(th) > 0 {
		t.Errorf("Sqrt(%v): %v", x, actual)
	}
}
