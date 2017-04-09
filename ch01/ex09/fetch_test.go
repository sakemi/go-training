package main

import "testing"

func TestInsertPrefixIfNeeded(t *testing.T) {
	actual := insertPrefixIfNeeded("abcde", "http://")
	expect := "http://abcde"
	if actual != expect {
		t.Error(expect, actual)
	}

	actual = insertPrefixIfNeeded("http://abcde", "http://")
	expect = "http://abcde"
	if actual != expect {
		t.Error(expect, actual)
	}
}
