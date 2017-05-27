package main

import "fmt"

func main() {
	vals := []int{-50, -1, 0, 1, 50}
	fmt.Printf("Vals:%v\n", vals)

	m, _ := max1(vals...)
	fmt.Printf("max1:%d", m)
}

func max1(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("No value contains in the argument.")
	}
	max := max2(vals[0], vals[1:]...)
	return max, nil
}

func max2(val int, vals ...int) int {
	if len(vals) == 0 {
		return val
	}
	max := val
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}

func min1(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("No value contains in the argument.")
	}
	min := min2(vals[0], vals[1:]...)
	return min, nil
}

func min2(val int, vals ...int) int {
	if len(vals) == 0 {
		return val
	}
	min := val
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min
}
