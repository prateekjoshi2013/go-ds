package queue_test

import (
	"testing"

	"github.com/prateekjoshi2013/gods/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()
	if q.Size() != 0 {
		t.Error("empty queue not created")
	}
	if _, err := q.Dequeue(); err == nil {
		t.Error("empty queue not created")
	}
	q.Enqueue(1)
	element, err := q.Dequeue()
	if err != nil {
		t.Error("empty queue after enqueue")
	}
	if element != 1 {
		t.Error("wrong element dequed")
	}
	if q.Size() != 0 {
		t.Error("queue not empty after dequeue")
	}
}
