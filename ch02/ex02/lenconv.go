package main

import "fmt"

type Meters float64
type Feet float64

const constant float64 = 3.28

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }

func MToF(m Meters) Feet { return Feet(float64(m) * constant) }
func FToM(f Feet) Meters { return Meters(float64(f) / constant) }
