package main

import "testing"

func TestReverse(t *testing.T) {
	s := []string{"a", "世界", "abcd", "12345", "辦ab12"}
	expected := []string{"a", "界世", "dcba", "54321", "21ba辦"}
	var actual string
	for i, v := range s {
		b := []byte(v)
		reverse(b)
		actual = string(b)
		if actual != expected[i] {
			t.Errorf("expected:%v, actual:%v\n", expected[i], actual)
		}
	}
}

// Test if capacity does not increase.
func TestCapacity(t *testing.T) {
	// set up
	s := []string{"a", "世界", "abcd", "12345", "辦ab12"}
	var expected []int
	for _, v := range s {
		expected = append(expected, cap([]byte(v)))
	}

	// execute
	var actual []int
	for _, v := range s {
		reverse([]byte(v))
		actual = append(actual, cap([]byte(v)))
	}

	// test
	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("expected:%v, actual:%v\n", v, actual[i])
		}
	}

}
