package main

import "math"

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(round(float64(c*9/5+32), .5, 2))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(round(float64((f-32)*5/9), .5, 2))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(round(float64(c-AbsoluteZeroC), .5, 2))
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
