package main

import "testing"

const testLen = 12

//testset contains 0-9, 100, 9999
//len=12
func newTestSet() *IntSet {
	set := IntSet{}
	for i := 0; i < 10; i++ {
		set.Add(i)
	}
	set.Add(100)
	set.Add(9999)
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
