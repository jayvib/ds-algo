package set

func New() *IntSet {
	return &IntSet{m: make(map[int]struct{})}
}

type IntSet struct {
	m map[int]struct{}
}

func (s *IntSet) IsExists(e int) bool {
	_, ok := s.m[e]
	return ok
}

func (s *IntSet) AddElement(e int) {
	if !s.IsExists(e) {
		s.m[e] = struct{}{}
	}
}

func (s *IntSet) DeleteElement(e int) {
	delete(s.m, e)
}

func (s *IntSet) Intersection(anotherSet *IntSet) *IntSet {
	intersectedSet := New()
	for k := range s.m {
		if anotherSet.IsExists(k) {
			intersectedSet.AddElement(k)
		}
	}
	return intersectedSet
}
