package stack_test

import (
	"testing"

	"github.com/prateekjoshi2013/go-ds/stack"
)

func TestStack(t *testing.T){
	s:=stack.New[int]()
	s.Push(1)
	s.Push(3)
	element,err:=s.Pop()
	if err!=nil{
		t.Error("stack empty after push")
	} 
	if element !=3{
		t.Error("wrong element popped")
	}
}
