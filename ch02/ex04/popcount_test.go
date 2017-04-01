package main

import "testing"

func TestShiftPopCount(t *testing.T) {
	actual := ShiftPopCount(0)
	if actual != 0 {
		t.Error(0, actual)
	}

	actual = ShiftPopCount(0xf)
	if actual != 4 {
		t.Error(4, actual)
	}

	actual = ShiftPopCount(^uint64(0))
	if actual != 64 {
		t.Error(64, actual)
	}
}

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
