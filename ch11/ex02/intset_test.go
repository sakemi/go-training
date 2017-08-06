package main

import "testing"

func TestHas(t *testing.T) {
	tests := []struct {
		intSet IntSet
		set    map[int]bool
	}{
		{IntSet{[]uint64{1}}, map[int]bool{0: true}},
		{IntSet{[]uint64{3, 1}}, map[int]bool{0: true, 1: true, 64: true}},
	}
	for _, test := range tests {
		for i := 0; i < 100; i++ {
			has := test.intSet.Has(i)
			_, ok := test.set[i]
			if has != ok {
				t.Errorf("set:%v, intSet:%v", test.set, test.intSet)
			}
		}
	}

}

func TestAdd(t *testing.T) {
	tests := [][]int{
		[]int{0},
		[]int{1, 2, 3},
		[]int{1, 1},
		[]int{0, 9999999},
	}

	for _, test := range tests {
		set := map[int]bool{}
		intSet := IntSet{}
		for _, v := range test {
			set[v] = true
			intSet.Add(v)
		}

		if !equals(set, intSet) {
			t.Errorf("set:%v, intSet:%v", set, intSet)
		}
	}
}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		intSet1 IntSet
		intSet2 IntSet
		set1    map[int]bool
		set2    map[int]bool
	}{
		{IntSet{[]uint64{1}}, IntSet{[]uint64{2}}, map[int]bool{0: true}, map[int]bool{1: true}},
		{IntSet{[]uint64{3, 1}}, IntSet{[]uint64{8, 2}}, map[int]bool{0: true, 1: true, 64: true}, map[int]bool{3: true, 65: true}},
	}

	for _, test := range tests {
		test.intSet1.UnionWith(&test.intSet2)
		unionedSet := test.set1
		for k, v := range test.set2 {
			unionedSet[k] = v
		}

		if !equals(unionedSet, test.intSet1) {
			t.Error(test)
		}
	}
}

func equals(set map[int]bool, intSet IntSet) bool {
	size := 0
	for _, v := range intSet.words {
		size += PopCount(v)
	}
	if len(set) != size {
		return false
	}

	for val := range set {
		if !intSet.Has(val) {
			return false
		}
	}

	return true
}
