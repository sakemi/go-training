package main

import "fmt"

func main() {
	strings := []string{"a", "b", "b", "aaa", "bb", "bb"}
	fmt.Printf("%v\n", dup(strings))
}

func dup(strings []string) []string {
	prev := ""
	i := 0
	for _, s := range strings {
		if i == 0 && prev == "" { //strings[0]が空文字のとき
			strings[i] = s
			i++
		} else if prev != s {
			strings[i] = s
			i++
		}
		prev = s
	}
	return strings[:i]
}
