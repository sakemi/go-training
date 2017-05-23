package main

import "testing"

const trial = 100

func TestTopoSort(t *testing.T) {
	for i := 0; i < trial; i++ {
		order := topoSort(prereqs)
		if key, ok := checkOrder(order); !ok {
			t.Errorf("Cannot take %s", key)
		}
	}
}

func checkOrder(order []string) (string, bool) {
	seen := map[string]bool{}
	for _, v := range order {
		if prereqs[v] == nil {
			seen[v] = true
			continue
		}
		for k := range prereqs[v] {
			if seen[k] == false {
				return v, false
			}
		}
		seen[v] = true
	}
	return "", true
}
