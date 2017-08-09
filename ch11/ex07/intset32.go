package main

import (
	"bytes"
	"fmt"
)

type IntSet32 struct {
	words []uint32
}

func (s *IntSet32) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet32) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet32) UnionWith(t *IntSet32) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet32) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet32) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount32(word)
	}
	return count
}

func (s *IntSet32) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/32, uint(x%32)
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet32) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet32) Copy() *IntSet32 {
	set := IntSet32{}
	set.words = make([]uint32, len(s.words))
	copy(set.words, s.words)
	return &set
}

func (s *IntSet32) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

func (s *IntSet32) IntersectWith(t *IntSet32) {
	for i, v := range t.words {
		if i < len(s.words) {
			s.words[i] &= v
		}
	}
	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet32) DifferenceWith(t *IntSet32) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			return
		}
	}
}

func (s *IntSet32) SymmetricDifferenceWith(t *IntSet32) {
	t2 := t.Copy()
	t2.IntersectWith(s)
	s.UnionWith(t)
	s.DifferenceWith(t2)
}

func (s *IntSet32) Elem() []uint32 {
	elem := []uint32{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				elem = append(elem, uint32(32*i+j))
			}
		}
	}
	return elem
}

func popCount32(x uint32) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
