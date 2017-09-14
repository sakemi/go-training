package main

import "testing"

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

type test struct {
	input *data
	want  string
}

const base = "http://hoge/fuga"

func TestPack(t *testing.T) {
	testCase := []test{
		test{
			&data{
				[]string{"l1", "l2", "l3"},
				10,
				true,
			},
			"l=l1&l=l2&l=l3&max=10&x=true",
		},
		test{
			&data{},
			"max=0&x=false",
		},
	}

	for _, tc := range testCase {
		got := Pack(tc.input).RawQuery
		if got != tc.want {
			t.Errorf("Pack(%v)=%s want:%s", tc.input, got, tc.want)
		}
	}
}
