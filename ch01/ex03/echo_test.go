package ex03

import (
	"strings"
	"testing"
)

var multiArgs [10]string = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
var singleArg [1]string = [1]string{"a"}

func BenchmarkJoinSingleArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(singleArg[:], " ")
	}
}

func BenchmarkAddOperatorSingleArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range singleArg[:] {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkJoinMultiArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(multiArgs[:], " ")
	}
}

func BenchmarkAddOperatorMultiArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range multiArgs[:] {
			s += sep + arg
			sep = " "
		}
	}
}
