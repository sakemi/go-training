package main

import (
	"fmt"
	"testing"
)

// for windows
func TestRunEditor(t *testing.T) {
	str, err := runEditor("C:\\WINDOWS\\system32\\notepad.exe")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(str)
}
