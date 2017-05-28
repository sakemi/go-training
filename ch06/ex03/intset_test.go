package main

import "testing"

const testLen = 10

//testset contains 0-9
func newTestSet() *IntSet {
	set := IntSet{}
	for i := 0; i < 10; i++ {
		set.Add(i)
	}
	return &set
}

func TestLen(t *testing.T) {
	test := newTestSet()
	if l := test.Len(); l != testLen {
		t.Errorf("Expected %d but actual %d", testLen, l)
	}

	test.Add(10)
	if l := test.Len(); l != testLen+1 {
		t.Errorf("Expected %d but actual %d", testLen+1, l)
	}
}

func TestRemove(t *testing.T) {
	test := newTestSet()
	test.Remove(0)
	if test.Has(0) {
		t.Errorf("Failed to remove.")
	}
	if l := test.Len(); l != testLen-1 {
		t.Errorf("Expected %d but actual %d", testLen-1, l)
	}
}

func TestClear(t *testing.T) {
	test := newTestSet()
	test.Clear()
	if l := test.Len(); l != 0 {
		t.Errorf("%d elements are left", l)
	}
}

func TestCopy(t *testing.T) {
	test := newTestSet()
	c := test.Copy()
	if &test.words == &c.words {
		t.Errorf("Copy and source refer to same slice.")
	}
	if !equals(test.words, c.words) {
		t.Errorf("Failed to copy elements.")
	}
}

func TestAddAll(t *testing.T) {
	test := newTestSet()
	x := []int{10, 11, 12}
	test.AddAll(x...)

	expected := newTestSet()
	for _, v := range x {
		expected.Add(v)
	}
	if !equals(expected.words, test.words) {
		t.Errorf("Expect %v but actual %v", expected, test)
	}
}

func TestIntersectWith(t *testing.T) {
	s1 := newTestSet()
	s2 := new(IntSet)
	s2.AddAll(1, 2, 3, 10, 11, 12)
	expected := new(IntSet)
	expected.AddAll(1, 2, 3)

	s1.IntersectWith(s2)
	if !equals(expected.words, s1.words) {
		t.Errorf("Expect %v but actual %v", expected, s1)
	}
}

func TestDifferenceWith(t *testing.T) {
	s1 := newTestSet()
	s2 := new(IntSet)
	s2.AddAll(1, 2, 3, 10, 11, 12)
	expected := new(IntSet)
	expected.AddAll(0, 4, 5, 6, 7, 8, 9)

	s1.DifferenceWith(s2)
	if !equals(expected.words, s1.words) {
		t.Errorf("Expect %v but actual %v", expected, s1)
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	s1 := newTestSet()
	s2 := new(IntSet)
	s2.AddAll(1, 2, 3, 10, 11, 12)
	expected := new(IntSet)
	expected.AddAll(0, 4, 5, 6, 7, 8, 9, 10, 11, 12)

	s1.SymmetricDifferenceWith(s2)
	if !equals(expected.words, s1.words) {
		t.Errorf("Expect %v but actual %v", expected, s1)
	}
}

func equals(x, y []uint64) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
