package main

import (
	"runtime"
	"testing"
)

func Benchmark1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark4(b *testing.B) {
	runtime.GOMAXPROCS(4)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark8(b *testing.B) {
	runtime.GOMAXPROCS(8)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark16(b *testing.B) {
	runtime.GOMAXPROCS(16)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark32(b *testing.B) {
	runtime.GOMAXPROCS(32)
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}
