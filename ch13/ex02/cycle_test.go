package main

import "testing"

type link struct {
	value string
	tail  *link
}

type test struct {
	input *link
	want  bool
}

func TestIsCycle(t *testing.T) {
	// cycle
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	// not cycle
	d, e := &link{value: "d"}, &link{value: "e"}
	d.tail, e.tail = e, nil

	testCase := []test{
		test{a, true},
		test{b, true},
		test{c, true},
		test{d, false},
		test{e, false},
	}

	for _, tc := range testCase {
		if got := IsCycle(tc.input); tc.want != got {
			t.Errorf("IsCycle(%v)=%v", tc.input, got)
		}
	}
}
