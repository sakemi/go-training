package main

import "testing"

func BenchmarkPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
	}
}

func BenchmarkShiftPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(0)
	}
}

func BenchmarkClearPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(0)
	}
}

func BenchmarkPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(^uint64(0))
	}
}

func BenchmarkShiftPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(^uint64(0))
	}
}

func BenchmarkClearPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(^uint64(0))
	}
}
