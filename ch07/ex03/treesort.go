package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	t := new(tree)
	vals := []int{1, 3, 5, 6, 3, 6, 87, 3462, 7, 6, 653}
	for _, v := range vals {
		t = add(t, v)
	}
	fmt.Println(t)
}

func (t *tree) String() string {
	vals := []string{}
	var walk func(t *tree)
	walk = func(t *tree) {
		if t.left != nil {
			walk(t.left)
		}
		vals = append(vals, strconv.Itoa(t.value))
		if t.right != nil {
			walk(t.right)
		}
	}

	walk(t)
	return strings.Join(vals, " ")
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
