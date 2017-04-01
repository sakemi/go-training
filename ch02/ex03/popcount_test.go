package main

import "testing"

func TestLoopPopCount(t *testing.T) {
	actual := LoopPopCount(0)
	if actual != 0 {
		t.Error(0, actual)
	}

	actual = LoopPopCount(0xf)
	if actual != 4 {
		t.Error(4, actual)
	}

	actual = LoopPopCount(^uint64(0))
	if actual != 64 {
		t.Error(64, actual)
	}
}

func BenchmarkPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
	}
}

func BenchmarkLoopPopCountZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(0)
	}
}

func BenchmarkPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(^uint64(0))
	}
}

func BenchmarkLoopPopCountMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(^uint64(0))
	}
}
