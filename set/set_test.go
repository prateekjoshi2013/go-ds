package set_test

import (
	"testing"

	"github.com/prateekjoshi2013/gods/set"
)

func TestSet(t *testing.T) {
	u := set.New[int]()
	v := set.New[int]()
	w := set.New[int]()

	if u.Size() != 0 && v.Size() != 0 && w.Size() != 0 {
		t.Fatal("non empty set created")
	}

	u.Add(1)
	u.Add(2)
	u.Add(3)
	v.Add(3)
	v.Add(4)
	v.Add(5)
	w.Add(5)
	w.Add(6)
	w.Add(7)
	if !u.Has(1) || !v.Has(3) || !w.Has(6) {
		t.Error("set does not have element which was added")
	}

	uAndV := u.Intersection(v)
	if uAndV.Size() != 1 || !uAndV.Has(3) {
		t.Error("intersection is not working as expected")
	}

	vAndW := v.Union(w)
	if vAndW.Size() != 5 {
		t.Error("union is not working as expected")
	}
}
