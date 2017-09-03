package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type data struct {
	I int
	F float64
	B bool
	A []int
	M map[string]int
	S struct {
		X   int
		Str string
	}
}

func TestMarshal(t *testing.T) {
	test := data{
		I: 1,
		F: 3.14,
		B: true,
		A: []int{1, 2, 3},
		M: map[string]int{
			"hoge": 1,
			"fuga": 2,
		},
		S: struct {
			X   int
			Str string
		}{
			X:   1,
			Str: "foo",
		},
	}

	b, err := Marshal(test)
	if err != nil {
		t.Error(err)
	}

	got := data{}
	if err := json.Unmarshal(b, &got); err != nil {
		fmt.Fprintln(os.Stderr, string(b))
		t.Error(err)
	}

	b, err = json.Marshal(test)
	if err != nil {
		t.Error(err)
	}

	want := data{}
	if err := json.Unmarshal(b, &want); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(&got, &want) {
		t.Errorf("got:%v\nwant:%v", got, want)
	}
}
