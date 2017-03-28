package main

import "testing"

func TestKgToLb(t *testing.T) {
	if actual := KgToLb(1); actual != 2.204 {
		t.Error(Kg(1), actual)
	}
}

func TestKgToKan(t *testing.T) {
	if actual := KgToKan(15); actual != 4 {
		t.Error(Kg(15), actual)
	}
}

func TestLbToKg(t *testing.T) {
	if actual := LbToKg(2.204); actual != 1 {
		t.Error(Lb(2.204), actual)
	}
}

func TestLbToKan(t *testing.T) {
	if actual := LbToKan(33.06); actual != 4 {
		t.Error(Lb(33.06), actual)
	}
}

func TestKanToKg(t *testing.T) {
	if actual := KanToKg(4); actual != 15 {
		t.Error(Kan(4), actual)
	}
}

func TestKanToLb(t *testing.T) {
	if actual := KanToLb(4); actual != 33.06 {
		t.Error(Kan(4), actual)
	}
}
