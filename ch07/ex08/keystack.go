package main

import "fmt"

const (
	keyTitle  = "title"
	keyArtist = "artist"
	keyAlubm  = "alubm"
	keyYear   = "year"
	keyLength = "length"
)

var keys map[string]bool = map[string]bool{keyTitle: true, keyArtist: true, keyAlubm: true, keyYear: true, keyLength: true}

type keyStack struct {
	stack []string
}

func (s *keyStack) push(key string) {
	for i, v := range s.stack {
		if v == key {
			//romove
			s.stack = append(s.stack[:i], s.stack[i+1:]...)
			break
		}
	}
	s.stack = append(s.stack, key)
}

func (s *keyStack) pop() (key string, err error) {
	if len(s.stack) == 0 {
		key = ""
		err = fmt.Errorf("Empty")
		return
	}
	key = s.stack[len(s.stack)-1]
	err = nil
	return
}
