package linkedlist

import (
	"testing"
)


func TestQ21(t *testing.T) {
	l := LinkedList{}
	if &l == nil {
		t.Error("LinkedList Is Not Create")
	}

	l.AddFirst(INTVAL(1))
	l.AddFirst(INTVAL(2))
	l.AddFirst(INTVAL(3))
	
	if !l.head.value.equal(INTVAL(3)) || !l.head.next.value.equal(INTVAL(2)) || !l.head.next.next.value.equal(INTVAL(1)) {
		t.Error("Add Failure")
	}
}
