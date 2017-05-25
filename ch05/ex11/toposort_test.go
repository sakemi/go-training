package main

import "testing"

const trial = 100

var noCirculation = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TestTopoSort(t *testing.T) {
	for i := 0; i < trial; i++ {
		order, err := topoSort(noCirculation)
		if err != nil {
			t.Error(err)
		}
		if key, ok := checkOrder(order, noCirculation); !ok {
			t.Errorf("Cannot take %s", key)
		}
	}
}

func checkOrder(order []string, m map[string][]string) (string, bool) {
	seen := map[string]bool{}
	for _, v := range order {
		if m[v] == nil {
			seen[v] = true
			continue
		}
		for _, k := range m[v] {
			if seen[k] == false {
				return v, false
			}
		}
		seen[v] = true
	}
	return "", true
}

var circulation1 = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var circulation2 = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math", "networks"}, //circulation!!
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TestCirculation(t *testing.T) {
	if _, err := topoSort(circulation1); err == nil {
		t.Error(circulation1)
	}
	if _, err := topoSort(circulation2); err == nil {
		t.Error(circulation2)
	}
}
