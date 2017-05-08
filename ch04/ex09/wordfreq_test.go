package main

import (
	"os"
	"testing"
)

func TestCountWords(t *testing.T) {
	f, err := os.Open("test.txt")
	defer f.Close()
	if err != nil {
		t.Error("failed to open file")
	}
	counts := make(map[string]int)
	countWords(f, counts)

	if counts["aaa"] != 1 {
		t.Errorf("\"aaa\" expected 1 but actual %d\n", counts["aaa"])
	}
	if counts["bbb"] != 2 {
		t.Errorf("\"bbb\" expected 2 but actual %d\n", counts["bbb"])
	}
	if counts["ccc"] != 3 {
		t.Errorf("\"ccc\" expected 3 but actual %d\n", counts["ccc"])
	}
	if counts["zzz"] != 0 {
		t.Errorf("\"zzz\" expected 0 but actual %d\n", counts["zzz"])
	}
}
