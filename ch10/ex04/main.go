package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type deps struct {
	Deps []string
}

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	in := os.Args[1]
	result := dependencies(in)
	for dep := range result {
		fmt.Println(dep)
	}
}

func dependencies(pkg string) map[string]bool {
	firstDeps := list(pkg)
	result := map[string]bool{}
	for _, dep := range firstDeps.Deps {
		result[dep] = true

		secondDeps := list(dep)
		for _, dep := range secondDeps.Deps {
			result[dep] = true
		}
	}
	return result
}

func list(pkg string) *deps {
	out, err := exec.Command("go", "list", "-json", pkg).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	depsEntity := new(deps)
	if err := json.Unmarshal(out, depsEntity); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	return depsEntity
}
