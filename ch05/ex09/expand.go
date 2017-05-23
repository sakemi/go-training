package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a $bbb $あああ aa$aa"
	e := expand(s, func(str string) string {
		return strings.Join([]string{"!!", str, "!!"}, "")
	})
	fmt.Println(e)
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if word[0] == '$' {
			words[i] = f(words[i][1:])
		}
	}
	return strings.Join(words, " ")
}
