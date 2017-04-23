package main

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "animation"
	s2 := "noitamina"
	if !isAnagram(s1, s2) {
		t.Errorf("%v, %v\n", s1, s2)
	}

	s1 = "THE COUNTRY SIDE"
	s2 = "NO CITY DUST HERE"
	if !isAnagram(s1, s2) {
		t.Errorf("%v, %v\n", s1, s2)
	}

	s1 = "かとう あい" //加藤 あい
	s2 = "あとう かい" //阿藤 快
	if !isAnagram(s1, s2) {
		t.Errorf("%v, %v\n", s1, s2)
	}

	s1 = "this is not anagram"
	s2 = "xxxxxxxxx"
	if isAnagram(s1, s2) {
		t.Errorf("%v, %v\n", s1, s2)
	}
}
