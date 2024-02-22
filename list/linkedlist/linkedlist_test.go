package linkedlist_test

import (
	"testing"

	"github.com/prateekjoshi2013/gods/list/linkedlist"
)

func TestLinkedList(t *testing.T) {
	l := linkedlist.New[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Print()
	l.Delete(1)
	l.Print()
	l.Delete(3)
	l.Print()
	l.Delete(4)
	l.Print()
	l.Delete(2)
	l.Print()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Print()
}
