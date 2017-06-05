package main

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%v)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%v%c%v)", b.x, b.op, b.y)
}

func (c call) String() string {
	s := []string{c.fn, "("}
	for i, arg := range c.args {
		if i > 0 {
			s = append(s, ", ", arg.String())
		}
		s = append(s, arg.String())
	}
	s = append(s, ")")
	return strings.Join(s, "")
}
