package main

type MapSet struct {
	set map[int]bool
}

func (s *MapSet) Has(x int) bool {
	_, ok := s.set[x]
	return ok
}

func (s *MapSet) Add(x int) {
	s.set[x] = true
}

func (s *MapSet) UnionWith(m *MapSet) {
	for k, _ := range m.set {
		s.set[k] = true
	}
}

func (s *MapSet) Len() int {
	cnt := 0
	for _, v := range s.set {
		if v {
			cnt++
		}
	}
	return cnt
}

func (s *MapSet) Remove(x int) {
	s.set[x] = false
}

func (s *MapSet) Clear() {
	for x, _ := range s.set {
		s.set[x] = false
	}
}

func (s *MapSet) Copy() *MapSet {
	m := map[int]bool{}
	for k, v := range s.set {
		m[k] = v
	}
	return &MapSet{m}
}

func (s *MapSet) AddAll(x ...int) {
	for _, v := range x {
		s.set[v] = true
	}
}

func (s *MapSet) IntersectWith(m *MapSet) {
	for k, _ := range s.set {
		v, ok := m.set[k]
		if !ok || !v {
			s.set[k] = false
		}
	}
}

func (s *MapSet) DifferenceWith(m *MapSet) {
	for k, _ := range s.set {
		if m.set[k] {
			s.set[k] = false
		}
	}
}

func (s *MapSet) SymmetricDifferenceWith(m *MapSet) {
	m2 := m.Copy()
	m2.IntersectWith(s)
	s.UnionWith(m)
	s.DifferenceWith(m2)
}
