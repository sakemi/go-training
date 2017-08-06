package main

import (
	"math/rand"
	"testing"
)

var intSet1 IntSet
var intSet2 IntSet
var mapSet1 MapSet
var mapSet2 MapSet

func init() {
	rng := newRng0()
	intSet1 = IntSet{}
	mapSet1 = MapSet{map[int]bool{}}
	for i := 0; i < 1000; i++ {
		x := rng.Intn(10000)
		intSet1.Add(x)
		mapSet1.Add(x)
	}

	intSet2 = IntSet{}
	mapSet2 = MapSet{map[int]bool{}}
	for i := 0; i < 1000; i++ {
		x := rng.Intn(10000)
		intSet2.Add(x)
		mapSet2.Add(x)
	}
}

func BenchmarkHas(b *testing.B) {
	rng := newRng0()
	for i := 0; i < b.N; i++ {
		intSet1.Has(rng.Intn(10000))
	}
}

func BenchmarkHasByMap(b *testing.B) {
	rng := newRng0()
	for i := 0; i < b.N; i++ {
		mapSet1.Has(rng.Intn(10000))
	}
}

func BenchmarkAdd10000(b *testing.B) {
	rng := newRng0()
	intSet := IntSet{}
	for i := 0; i < b.N; i++ {
		intSet.Add(rng.Intn(10000))
	}
}

func BenchmarkAddByMap10000(b *testing.B) {
	rng := newRng0()
	set := MapSet{map[int]bool{}}
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(10000))
	}
}

func BenchmarkAdd10(b *testing.B) {
	rng := newRng0()
	intSet := IntSet{}
	for i := 0; i < b.N; i++ {
		intSet.Add(rng.Intn(10))
	}
}

func BenchmarkAddByMap10(b *testing.B) {
	rng := newRng0()
	set := MapSet{map[int]bool{}}
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(10000))
	}
}

func BenchmarkUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSet1.UnionWith(&intSet2)
	}
}

func BenchmarkUnionWithByMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapSet1.UnionWith(&mapSet2)
	}
}

func newRng0() *rand.Rand {
	return rand.New(rand.NewSource(0))
}
