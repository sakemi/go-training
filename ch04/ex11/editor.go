package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func runEditor(editor string) (string, error) {
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	name := tmp.Name()
	tmp.Close()
	cmd := exec.Command(editor, name)
	if err := cmd.Run(); err != nil {
		return "", err
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	if err := os.Remove(name); err != nil {
		return "", err
	}
	return string(data), nil
}
