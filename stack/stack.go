package stack

import "errors"

type Stack[T any] struct {
	elements []T
}

func New[T any]() Stack[T] {
	return Stack[T]{[]T{}}
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var t T
		return t, errors.New("stack is empty")
	}
	element := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]
	return element, nil
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var t T
		return t, errors.New("empty stack")
	}
	element := s.elements[s.Size()-1]
	return element, nil
}
