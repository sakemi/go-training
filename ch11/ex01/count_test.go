package main

import (
	"bufio"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{"a", map[rune]int{'a': 1}, [utf8.UTFMax + 1]int{0, 1, 0, 0, 0}, 0},
		{"abbccc", map[rune]int{'a': 1, 'b': 2, 'c': 3}, [utf8.UTFMax + 1]int{0, 6, 0, 0, 0}, 0},
		{"\xf0", map[rune]int{}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}, 1},
		{"a©あ𠀋", map[rune]int{'a': 1, '©': 1, 'あ': 1, '𠀋': 1}, [utf8.UTFMax + 1]int{0, 1, 1, 1, 1}, 0},
		{"a©あ𠀋　\n a\xf0", map[rune]int{'a': 2, '©': 1, 'あ': 1, '𠀋': 1, '　': 1, '\n': 1, ' ': 1}, [utf8.UTFMax + 1]int{0, 4, 1, 2, 1}, 1},
	}

	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		counts := make(map[rune]int)
		var utflen [utf8.UTFMax + 1]int
		invalid := 0

		count(in, counts, &utflen, &invalid)
		checkCounts(t, test.input, counts, test.counts)
		checkUTFlen(t, test.input, utflen, test.utflen)
		checkInvalid(t, test.input, invalid, test.invalid)
	}
}

func checkCounts(t *testing.T, input string, got, want map[rune]int) {
	if len(got) != len(want) {
		t.Errorf("input: %s, got:%v, want%v", input, got, want)
		return
	}
	for k, v := range got {
		if v != want[k] {
			t.Errorf("input: %s, got:%v, want:%v", input, got, want)
			return
		}
	}
}

func checkUTFlen(t *testing.T, input string, got, want [utf8.UTFMax + 1]int) {
	for i := 0; i < utf8.UTFMax+1; i++ {
		if got[i] != want[i] {
			t.Errorf("input: %s, got:%v, want:%v", input, got, want)
			return
		}
	}
}

func checkInvalid(t *testing.T, input string, got, want int) {
	if got != want {
		t.Errorf("input: %s, got:%d, want:%d", input, got, want)
		return
	}
}
