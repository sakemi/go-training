package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

type data struct {
	Label   string `http:"l"`
	Mail    string `email:"m"`
	Card    string `card:"c"`
	ZipCode int    `zipcode:"z"`
}

type params struct {
	label   string
	mail    string
	card    string
	zipCode string
}

func TestUnpackValidation(t *testing.T) {
	errorCase := []params{
		params{
			"label",
			"malformedAddress",
			"0123456789012345",
			"01234",
		},
		params{
			"label",
			"foo@barmail.com",
			"malformed",
			"01234",
		},
		params{
			"label",
			"foo@barmail.com",
			"0123456789012345",
			"malformed",
		},
	}

	for _, ec := range errorCase {
		url := createURL(ec.label, ec.mail, ec.card, ec.zipCode)
		req, err := http.NewRequest("Get", url, strings.NewReader(""))
		if err != nil {
			t.Error(err)
		}
		d := &data{}
		err = Unpack(req, d)
		if err == nil {
			t.Error("Failed to validate.")
		}
	}
}

func TestPack(t *testing.T) {
	testCase := []params{
		params{
			"label",
			"foo@barmail.com",
			"0123456789012345",
			"01234",
		},
	}

	for _, tc := range testCase {
		url := createURL(tc.label, tc.mail, tc.card, tc.zipCode)
		req, err := http.NewRequest("Get", url, strings.NewReader(""))
		if err != nil {
			t.Error(err)
		}
		d := &data{}
		err = Unpack(req, d)
		if err != nil {
			t.Error(err)
		}

		if d.Label != tc.label {
			t.Errorf("Failed to Unpack: want %v, got %v", tc.label, d.Label)
		}
		if d.Mail != tc.mail {
			t.Errorf("Failed to Unpack: want %v, got %v", tc.mail, d.Mail)
		}
		if d.Card != tc.card {
			t.Errorf("Failed to Unpack: want %v, got %v", tc.card, d.Card)
		}
		zipcode, _ := strconv.Atoi(tc.zipCode)
		if d.ZipCode != zipcode {
			t.Errorf("Failed to Unpack: want %v, got %v", zipcode, d.ZipCode)
		}
	}
}

func createURL(l, m, c, z string) string {
	return fmt.Sprintf("http://hoge.com/?l=%s&m=%s&c=%s&z=%s", l, m, c, z)
}
