package main

import "testing"

func TestBundleSpace(t *testing.T) {
	s := []string{"a", "a bc d", "a    b  cd e　 f", "  abcdef  "}
	expected := []string{"a", "a bc d", "a b cd e f", " abcdef "}
	var actual string
	for i, v := range s {
		b := []byte(v)
		actual = string(bundleSpace(b))
		if actual != expected[i] {
			t.Errorf("input: %v, expected:%v, actual:%v\n", s[i], expected[i], actual)
		}
	}
}
