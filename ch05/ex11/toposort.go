package main

import (
	"fmt"
	"os"
	"sort"
)

var prereqs = map[string][]string{
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

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	var walk func(items []string) error

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	checkCirculate := func(root string) error {
		walk = func(items []string) error {
			for _, item := range items {
				if item == root {
					return fmt.Errorf("Circulation Error: %s", root)
				}
				if err := walk(m[item]); err != nil {
					return err
				}
			}
			return nil
		}

		if err := walk(m[root]); err != nil {
			return err
		}
		return nil
	}

	var keys []string
	for key := range m {
		if err := checkCirculate(key); err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order, nil
}
