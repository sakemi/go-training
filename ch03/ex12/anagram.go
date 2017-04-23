package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "animation"
	s2 := "noitamina"
	fmt.Printf("%v, %v %v\n", s1, s2, isAnagram("animation", "noitamina"))
}

func isAnagram(s1, s2 string) bool {
	//remove space
	r1 := []rune(spaceMap(s1))
	r2 := []rune(spaceMap(s2))
	for i := 0; i < len(r1); i++ {
		find := false
		for j := 0; j < len(r2); j++ {
			if r1[i] == r2[j] {
				r2 = append(r2[:j], r2[j+1:]...)
				find = true
				break
			}
		}
		if !find {
			break
		}
	}

	if len(r2) == 0 {
		return true
	}
	return false
}

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
