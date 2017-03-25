package main

import "testing"

const (
	httpPrefix = "http://"
	url        = "hoge/fuga"
)

func TestInsertPrefixIfNeeded(t *testing.T) {
	//setup
	httpUrl := httpPrefix + url
	noHttpUrl := url

	//execute
	actual1 := insertPrefixIfNeeded(httpUrl, httpPrefix)
	actual2 := insertPrefixIfNeeded(noHttpUrl, httpPrefix)

	//test
	expected := httpPrefix + url
	if actual1 != expected {
		t.Error("URL with http:", actual1)
	}
	if actual2 != expected {
		t.Error("URL without http", actual2)
	}
}
