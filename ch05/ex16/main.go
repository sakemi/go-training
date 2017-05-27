package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("/", "foo", "bar", "baz"))
}

func join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
