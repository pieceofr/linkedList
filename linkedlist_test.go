package linkedlist

import (
	"testing"
)

func TestCreate(t *testing.T) {
	l := LinkedList{}
	if &l == nil {
		t.Error("LinkedList Is Not Create")
	}
}

func TestAddFirst(t *testing.T) {
	l := LinkedList{}
	if &l == nil {
		t.Error("LinkedList Is Not Create")
	}

	err := l.AddFirst(INTVAL(1))
	if err != nil {
		t.Error("Couldn't add value")
	}
	err = l.AddFirst(INTVAL(2))
	err = l.AddFirst(INTVAL(3))

	if !l.head.value.equal(INTVAL(3)) || !l.head.next.value.equal(INTVAL(2)) || !l.head.next.next.value.equal(INTVAL(1)) {
		t.Error("Add Failure")
	}
}

func TestTraverse(t *testing.T) {
	l := LinkedList{}
	c, node := l.Traverse()

	if c != 0 || node != nil {
		t.Error("count should be 0 and last node should be nil")
	}

	err := l.AddFirst(INTVAL(1))
	if err != nil {
		t.Error("Couldn't add value")
	}

	err = l.AddFirst(INTVAL(2))
	err = l.AddFirst(INTVAL(3))

	c, node = l.Traverse()

	if c != 3 || node == nil {
		t.Error("count should be 3 and last node should not be nil")
	}
}

func TestAddLast(t *testing.T) {
	l := LinkedList{}
	if &l == nil {
		t.Error("LinkedList Is Not Create")
	}

	err := l.AddLast(INTVAL(1))
	if err != nil {
		t.Error("Couldn't add value")
	}

	err = l.AddLast(INTVAL(2))
	err = l.AddLast(INTVAL(3))

	if !l.head.value.equal(INTVAL(1)) || !l.head.next.value.equal(INTVAL(2)) || !l.head.next.next.value.equal(INTVAL(3)) {
		t.Error("Add Failure")
	}
}
func TestAddAfter(t *testing.T) {
	l := LinkedList{}
	if &l == nil {
		t.Error("LinkedList Is Not Create")
	}

	err := l.AddAfter(INTVAL(1), INTVAL(2))

	if err == nil {
		t.Error("AddAfter Should Not Be Insert")
	}

	err = l.AddLast(INTVAL(1))
	err = l.AddAfter(INTVAL(1), INTVAL(2))
	if err != nil {
		t.Error("AddAfter Should Not Be Any Error")
	}
	if !l.head.value.equal(INTVAL(1)) || !l.head.next.value.equal(INTVAL(2)) {
		t.Error("AddAfter Failure")
	}
	err = l.AddAfter(INTVAL(1), INTVAL(3))

	if !l.head.value.equal(INTVAL(1)) || !l.head.next.value.equal(INTVAL(3)) || !l.head.next.next.value.equal(INTVAL(2)) {
		t.Error("AddAfter Failure")
	}
	//	l.PrintList()
}

func TestAddBefore(t *testing.T) {
	l := LinkedList{}
	err := l.AddBefore(INTVAL(1), INTVAL(2))

	if err == nil {
		t.Error("AddBefore Should Not Be Insert")
	}

	err = l.AddLast(INTVAL(1))
	err = l.AddBefore(INTVAL(1), INTVAL(2))
	if err != nil {
		t.Error("AddBefore Should Not Be Any Error")
	}

	if !l.head.value.equal(INTVAL(2)) || !l.head.next.value.equal(INTVAL(1)) {
		t.Error("AddBefore Failure")
	}
	err = l.AddBefore(INTVAL(1), INTVAL(3))

	if !l.head.value.equal(INTVAL(2)) || !l.head.next.value.equal(INTVAL(3)) || !l.head.next.next.value.equal(INTVAL(1)) {
		t.Error("AddBefore Failure")
	}
	//l.PrintList()
}

func TestDelete(t *testing.T) {
	l := LinkedList{}
	err := l.Delete(STRVAL("aa"))
	if err == nil {
		t.Error("Delete Operation Error")
	}
	err = l.AddLast(STRVAL("aa"))
	err = l.AddLast(STRVAL("bb"))
	err = l.AddLast(STRVAL("cc"))
	err = l.Delete(STRVAL("bb"))

	if !l.head.value.equal(STRVAL("aa")) || !l.head.next.value.equal(STRVAL("cc")) || l.head.next.next != nil {
		t.Error("AddBefore Failure")
	}
	err = l.Delete(STRVAL("cc"))
	if !l.head.value.equal(STRVAL("aa")) || l.head.next != nil {
		t.Error("AddBefore Failure")
	}
	err = l.Delete(STRVAL("aa"))
	if l.head != nil {
		t.Error("AddBefore Failure")
	}
}

func TestIsCircular(t *testing.T) {
	l := LinkedList{}
	circular, err := l.IsCircular()

	if err != nil || circular == true {
		t.Fatal("Fatal Error In Test Code")
	}

	err = l.AddLast(STRVAL("a"))

	if circular == true {
		t.Error("It Should Not Be Circular")
	}

	err = l.AddLast(STRVAL("b"))
	err = l.AddLast(STRVAL("c"))
	err = l.AddLast(STRVAL("d"))
	err = l.AddLast(STRVAL("e"))
	err = l.AddLast(STRVAL("f"))

	circular, err = l.IsCircular()
	if err != nil {
		t.Fatal("Fatal Error In Test Code")
	}
	if circular {
		t.Error("It Should Not Be Circular")
	}
	//Make circular
	_, tail := l.Traverse()
	tail.next = l.head.next.next // point to c

	circular, _ = l.IsCircular()
	if !circular {
		t.Error("the list should be circular")
	}
}

func TestNodeType(t *testing.T) {
	testint1 := INTVAL(1)
	testint2 := INTVAL(1)
	testint3 := INTVAL(2)

	teststr1 := STRVAL("aaa")
	teststr2 := STRVAL("aaa")
	teststr3 := STRVAL("bbb")

	testMyType1 := MyType{myvalue: 1, mystruct: Node{value: INTVAL(5)}}
	testMyType2 := MyType{myvalue: 1, mystruct: Node{value: INTVAL(5)}}
	testMyType3 := MyType{myvalue: 2, mystruct: Node{value: INTVAL(5)}}

	if !testint1.equal(testint2) {
		t.Error("Should be eqaul!")
	}
	if testint1.equal(testint3) {
		t.Error("Should be not eqaul!")
	}

	if !teststr1.equal(teststr2) {
		t.Error("Should be eqaul!")
	}
	if teststr1.equal(teststr3) {
		t.Error("Should be not eqaul!")
	}

	if !testMyType1.equal(testMyType2) {
		t.Error("Should be eqaul!")
	}
	if testMyType1.equal(testMyType3) {
		t.Error("Should be not eqaul!")
	}
}
