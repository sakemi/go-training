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
