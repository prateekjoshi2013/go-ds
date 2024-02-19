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
	element := s.elements[0]
	s.elements = s.elements[1:]
	return element, nil
}

func (s *Stack[T]) Push(element T) {
	s.elements = append([]T{element}, s.elements...)
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var t T
		return t, errors.New("empty stack")
	}
	element := s.elements[0]
	return element, nil
}
