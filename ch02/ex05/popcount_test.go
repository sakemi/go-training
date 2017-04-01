package main

import "testing"

func TestClearPopCount(t *testing.T) {
	actual := ClearPopCount(0)
	if actual != 0 {
		t.Error(0, actual)
	}

	actual = ClearPopCount(0xf)
	if actual != 4 {
		t.Error(4, actual)
	}

	actual = ClearPopCount(^uint64(0))
	if actual != 64 {
		t.Error(64, actual)
	}
}

func BenchmarkPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
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

func BenchmarkClearPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(^uint64(0))
	}
}
