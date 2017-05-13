package main

import "testing"

func TestSearch(t *testing.T) {
	index := []info{info{"t1", "u1"}, info{"t2", "u2"}}

	result := search("t1", index)
	if len(result) != 1 {
		t.Errorf("index, key, result: %v,%v,%v", index, "t1", result)
	}
	if result[0].Title != "t1" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t1", result)
	}
	if result[0].URL != "u1" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t1", result)
	}

	result = search("t2", index)
	if len(result) != 1 {
		t.Errorf("index, key, result: %v,%v,%v", index, "t2", result)
	}
	if result[0].Title != "t2" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t1", result)
	}
	if result[0].URL != "u2" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t2", result)
	}

	result = search("t", index)
	if len(result) != 2 {
		t.Errorf("index, key, result: %v,%v,%v", index, "t", result)
	}
	if result[0].Title != "t1" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t", result)
	}
	if result[0].URL != "u1" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t", result)
	}
	if result[1].Title != "t2" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t", result)
	}
	if result[1].URL != "u2" {
		t.Errorf("index, key, result: %v,%v,%v", index, "t", result)
	}
}
