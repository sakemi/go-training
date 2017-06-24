package main

import "testing"

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(1)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(2)
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(4)
	}
}

func Benchmark8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(8)
	}
}

func Benchmark16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(16)
	}
}

func Benchmark32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(32)
	}
}

func Benchmark64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(64)
	}
}

func Benchmark128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(128)
	}
}

func Benchmark256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(256)
	}
}

func Benchmark512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(512)
	}
}

func Benchmark1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(1024)
	}
}
