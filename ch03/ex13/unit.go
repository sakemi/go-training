package main

import (
	"fmt"
	"strconv"
)

const (
	_        = iota
	KiB unit = iota * 3
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

type unit int
type ByteSize float64

func main() {
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	fmt.Println(TiB)
	fmt.Println(PiB)
	fmt.Println(EiB)
	fmt.Println(ZiB)
	fmt.Println(YiB)
}

func (u unit) value() ByteSize {
	return exp(10, int(u))
}

func (u unit) String() string {
	return strconv.FormatFloat(float64(u.value()), 'e', -1, 64)
}

func exp(x, y int) (z ByteSize) {
	z = 1
	for i := 0; i < y; i++ {
		z *= ByteSize(x)
	}
	return
}
