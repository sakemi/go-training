package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

//Error ^uint(0) might be too small for shift of 63 (vet)
const platform = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/platform, uint(x%platform)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/platform, uint(x%platform)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platform; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", platform*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/platform, uint(x%platform)
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	set := IntSet{}
	set.words = make([]uint64, len(s.words))
	copy(set.words, s.words)
	return &set
}

func (s *IntSet) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	var l int
	if len(s.words) >= len(t.words) {
		l = len(s.words)
	} else {
		l = len(t.words)
	}

	for i := 0; i < l; i++ {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			return
		}
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	t2 := t.Copy()
	t2.IntersectWith(s)
	s.UnionWith(t)
	s.DifferenceWith(t2)
}

func (s *IntSet) Elem() []uint64 {
	elem := []uint64{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platform; j++ {
			if word&(1<<uint(j)) != 0 {
				elem = append(elem, uint64(platform*i+j))
			}
		}
	}
	return elem
}

func popCount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
