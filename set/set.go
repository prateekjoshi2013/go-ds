package set

type Set[T comparable] struct {
	elements map[T]bool
}

func New[T comparable]() Set[T] {
	return Set[T]{
		elements: make(map[T]bool),
	}
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Add(element T) bool {
	if _, ok := s.elements[element]; ok {
		return false
	} else {
		s.elements[element] = true
		return true
	}
}

func (s *Set[T]) Remove(element T) bool {
	if _, ok := s.elements[element]; ok {
		delete(s.elements, element)
		return true
	} else {
		return false
	}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	set := New[T]()
	for k := range s.elements {
		set.Add(k)
	}
	for k := range other.elements {
		set.Add(k)
	}
	return set
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	set := New[T]()
	for k := range s.elements {
		if other.Has(k) {
			set.Add(k)
		}
	}
	return set
}

func (s *Set[T]) Difference(other Set[T]) Set[T] {
	set := New[T]()
	for k := range s.elements {
		if !other.Has(k) {
			set.Add(k)
		}
	}
	return set
}

func (s *Set[T]) IsSubset(other Set[T]) bool {
	for k := range s.elements {
		if !other.Has(k) {
			return false
		}
	}
	return true
}
