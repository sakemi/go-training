package main

import (
	"math/rand"
	"testing"
	"time"
)

var intSet1 IntSet
var intSet2 IntSet
var int32Set1 IntSet32
var int32Set2 IntSet32
var mapSet1 MapSet
var mapSet2 MapSet

func initTestSet() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	intSet1 = IntSet{}
	int32Set1 = IntSet32{}
	mapSet1 = MapSet{map[int]bool{}}
	for i := 0; i < 1000; i++ {
		x := rng.Intn(10000)
		intSet1.Add(x)
		int32Set1.Add(x)
		mapSet1.Add(x)
	}

	intSet2 = IntSet{}
	int32Set2 = IntSet32{}
	mapSet2 = MapSet{map[int]bool{}}
	for i := 0; i < 1000; i++ {
		x := rng.Intn(10000)
		intSet2.Add(x)
		int32Set2.Add(x)
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

func BenchmarkAdd(b *testing.B) {
	rng := newRng0()
	intSet := IntSet{}
	for i := 0; i < b.N; i++ {
		intSet.Add(rng.Intn(10000))
	}
}

func BenchmarkAdd32(b *testing.B) {
	rng := newRng0()
	intSet := IntSet32{}
	for i := 0; i < b.N; i++ {
		intSet.Add(rng.Intn(10000))
	}
}

func BenchmarkAddByMap(b *testing.B) {
	rng := newRng0()
	set := MapSet{map[int]bool{}}
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(10000))
	}
}

func BenchmarkUnionWith(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.UnionWith(&intSet2)
	}
}

func BenchmarkUnionWith32(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		int32Set1.UnionWith(&int32Set2)
	}
}

func BenchmarkUnionWithByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.UnionWith(&mapSet2)
	}
}

func BenchmarkLen(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.Len()
	}
}

func BenchmarkLen32(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		int32Set1.Len()
	}
}

func BenchmarkLenByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.Len()
	}
}

func BenchmarkRemove(b *testing.B) {
	initTestSet()
	rang := newRng0()
	for i := 0; i < b.N; i++ {
		intSet1.Remove(rang.Intn(10000))
	}
}

func BenchmarkRemoveByMap(b *testing.B) {
	initTestSet()
	rang := newRng0()
	for i := 0; i < b.N; i++ {
		mapSet1.Remove(rang.Intn(10000))
	}
}

func BenchmarkClear(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.Clear()
	}
}

func BenchmarkClearByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.Clear()
	}
}

func BenchmarkCopy(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.Copy()
	}
}

func BenchmarkCopyByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.Copy()
	}
}

func BenchmarkAddAll(b *testing.B) {
	initTestSet()
	var x []int
	rang := newRng0()
	for i := 0; i < 1000; i++ {
		x = append(x, rang.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		intSet1.AddAll(x...)
	}
}

func BenchmarkAddAllByMap(b *testing.B) {
	initTestSet()
	var x []int
	rang := newRng0()
	for i := 0; i < 1000; i++ {
		x = append(x, rang.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		mapSet1.AddAll(x...)
	}
}

func BenchmarkIntersectWith(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.IntersectWith(&intSet2)
	}
}

func BenchmarkIntersectWithByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.IntersectWith(&mapSet2)
	}
}

func BenchmarkDifferenceWith(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.DifferenceWith(&intSet2)
	}
}

func BenchmarkDifferenceWithByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.DifferenceWith(&mapSet2)
	}
}

func BenchmarkSymmetricDifferenceWith(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		intSet1.SymmetricDifferenceWith(&intSet2)
	}
}

func BenchmarkSymmetricDifferenceWithByMap(b *testing.B) {
	initTestSet()
	for i := 0; i < b.N; i++ {
		mapSet1.SymmetricDifferenceWith(&mapSet2)
	}
}

func newRng0() *rand.Rand {
	return rand.New(rand.NewSource(0))
}
