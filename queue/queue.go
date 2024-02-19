package queue

import "errors"

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("empty queue")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var t T
		return t, errors.New("empty queue")
	}
	return q.elements[0], nil
}
func New[T any]() Queue[T] {
	return Queue[T]{[]T{}}
}
