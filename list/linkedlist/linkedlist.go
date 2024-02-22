package linkedlist

import "fmt"

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T comparable]() LinkedList[T] {
	return LinkedList[T]{nil, nil}
}

func (l *LinkedList[T]) Add(element T) {
	node := Node[T]{element, nil}
	if l.head == nil {
		l.head = &node
		l.tail = &node
		return
	} else {
		l.tail.next = &node
		l.tail = l.tail.next
		return
	}
}

func (l *LinkedList[T]) Insert(element T, pos int) {
	curr := l.head
	node := Node[T]{element, nil}
	for pos > 0 {
		curr = curr.next
		pos -= 1
	}
	if curr == l.tail {
		l.tail = &node
	}
	node.next = curr.next
	curr.next = &node

}

func (l *LinkedList[T]) Print() {
	curr := l.head
	for curr != nil {
		if curr == l.head {
			fmt.Printf("head -> %v", curr.val)
		} else if curr == l.tail {
			fmt.Printf(" -> %v <- tail ", curr.val)
		} else {
			fmt.Printf("-> %v", curr.val)
		}
		curr = curr.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Delete(element T) error {
	dummy := Node[T]{}
	dummy.next = l.head
	curr := &dummy
	for curr.next != nil && curr.next.val != element {
		curr = curr.next
	}
	if curr.next != nil {
		if curr.next == l.head && l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else if l.head == curr.next {
			l.head = curr.next.next
		} else if l.tail == curr.next {
			l.tail = curr.next
		}
		curr.next = curr.next.next
		return nil
	}
	return fmt.Errorf("element not found")
}

func (l *LinkedList[T]) Find(element T) (*Node[T], error) {
	curr := l.head
	for curr != nil {
		if curr.val == element {
			return curr, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}
