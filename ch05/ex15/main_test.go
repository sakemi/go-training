package main

import "testing"

type test struct {
	input    []int
	expected int
}

func TestMax1(t *testing.T) {
	testCase := []test{
		test{[]int{-100, 100, 0, 1, -3}, 100},
		test{[]int{1}, 1},
	}

	for _, tc := range testCase {
		actual, err := max1(tc.input...)
		if err != nil {
			t.Errorf("%v", err)
		}
		if actual != tc.expected {
			t.Errorf("max1(%v)=%d", tc.input, actual)
		}
	}
}

func TestMax1Error(t *testing.T) {
	testCase := []test{
		test{[]int{}, 0},
	}

	for _, tc := range testCase {
		_, err := max1(tc.input...)
		if err == nil {
			t.Errorf("Error is expected but no error.")
		}
	}
}

func TestMax2(t *testing.T) {
	testCase := []test{
		test{[]int{-100, 100, 0, 1, -3}, 100},
		test{[]int{1}, 1},
	}

	for _, tc := range testCase {
		actual := max2(tc.input[0], tc.input[1:]...)
		if actual != tc.expected {
			t.Errorf("max2(%v)=%d", tc.input, actual)
		}
	}
}

func TestMin1(t *testing.T) {
	testCase := []test{
		test{[]int{-100, 100, 0, 1, -3}, -100},
		test{[]int{1}, 1},
	}

	for _, tc := range testCase {
		actual, err := min1(tc.input...)
		if err != nil {
			t.Errorf("%v", err)
		}
		if actual != tc.expected {
			t.Errorf("min1(%v)=%d", tc.input, actual)
		}
	}
}

func TestMin1Error(t *testing.T) {
	testCase := []test{
		test{[]int{}, 0},
	}

	for _, tc := range testCase {
		_, err := min1(tc.input...)
		if err == nil {
			t.Errorf("Error is expected but no error.")
		}
	}
}

func TestMin2(t *testing.T) {
	testCase := []test{
		test{[]int{-100, 100, 0, 1, -3}, -100},
		test{[]int{1}, 1},
	}

	for _, tc := range testCase {
		actual := min2(tc.input[0], tc.input[1:]...)
		if actual != tc.expected {
			t.Errorf("min2(%v)=%d", tc.input, actual)
		}
	}
}
