package main

import "testing"

func TestShortNumber(t *testing.T) {
	s := "12"
	if actual := comma(s); actual != s {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestHead1(t *testing.T) {
	s := "1234"
	expected := "1,234"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestHead2(t *testing.T) {
	s := "12345"
	expected := "12,345"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestHead3(t *testing.T) {
	s := "123456"
	expected := "123,456"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestLongNumber(t *testing.T) {
	s := "1234567890"
	expected := "1,234,567,890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

//======== expanded function test ========

func TestSign(t *testing.T) {
	s := "+1234567890"
	expected := "+1,234,567,890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}

	s = "-1234567890"
	expected = "-1,234,567,890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestDecimal(t *testing.T) {
	s := "123456.7890"
	expected := "123,456.7890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}

	s = "1.234567890"
	expected = "1.234567890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}

func TestSignedDecimal(t *testing.T) {
	s := "+123456.7890"
	expected := "+123,456.7890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}

	s = "-1.234567890"
	expected = "-1.234567890"
	if actual := comma(s); actual != expected {
		t.Errorf("%v -> %v", s, actual)
	}
}
