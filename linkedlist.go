package linkedlist

import (
   "log"
   "errors"
)

type Node struct {
	next *Node
	value ListVal
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList)AddFirst(data ListVal) error {
	if l == nil {
		return errors.New("LinkedList is not initialized")
	}
	newNode := Node{value:data}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
	} else {
		l.head = &newNode
	}
	return nil
}

func (l *LinkedList)Traverse() (count int, tail *Node) {
	if l == nil || l.head == nil {
		return 0, nil
	}
	tail = l.head
	count = 1
	for tail.next != nil {
		tail = tail.next
		count ++ 
	}
	return count, tail
}

func (l *LinkedList)AddLast(data ListVal) error {
	if l == nil {
		return errors.New("LinkedList is not initialized")
	}
	newNode := Node{value:data}
	_, last := l.Traverse()
	if last !=nil {
		last.next = &newNode
	} else {
		l.head = &newNode // Head is nil
	}
	return nil
}

func (l *LinkedList)AddAfter(target ListVal, data ListVal) error {
	if l == nil {
		return errors.New("LinkedList is not initialized")
	}
	newNode := Node{value:data}
	iter := l.head
	for iter != nil  {
		if (iter.value.equal(target)) { // Found & insert
			newNode.next = iter.next
			iter.next = &newNode
			return nil
		}
		iter = iter.next
	}
	return errors.New("No Target Node Found")
}

func (l *LinkedList)AddBefore(target ListVal, data ListVal) error {
	if l == nil {
		return errors.New("LinkedList is not initialized")
	}
	newNode := Node{value:data}
	iter := l.head
	var pvnode *Node

	for iter != nil  {
		if (iter.value.equal(target)) { // Found & insert
			if pvnode != nil {
				pvnode.next = &newNode
				newNode.next = iter
			} else { // HEAD Node
				newNode.next = iter
				l.head = &newNode
			}
			return nil
		}
		pvnode = iter
		iter = iter.next
	}

	return errors.New("No Target Node Found")
}

func (l *LinkedList)Delete(target ListVal) error {
	if l == nil {
		return errors.New("LinkedList is not initialized")
	}

	iter := l.head
	var pvnode *Node

	for iter != nil  {
		if (iter.value.equal(target)) { // Found & insert
			if pvnode != nil {
				pvnode.next = iter.next	
			} else { // HEAD Node
				l.head = iter.next
			}
			return nil
		}
		pvnode = iter
		iter = iter.next
	}

	return errors.New("No Target Node Found")
}

func (l *LinkedList)IsCircular() (bool, error) {
	if l == nil {
		return false, errors.New("LinkedList is not initialized")
	}
	if l.head == nil {
		return false, nil
	}
	slow := l.head
	fast := l.head.next
	for fast !=nil && fast.next != nil { //slow != nil is not nessary since , fast will reach the end first
		if slow == fast {
			return true, nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false, nil
}


func (i *LinkedList)GetIterator() (ListIterator) {
	return ListIterator{nextNode:i.head}
}

func (l *LinkedList)PrintList() {
	if l == nil || l.head == nil {
		return
	}
	iter := l.head
	for  iter != nil {
		log.Println(iter.value)
		iter = iter.next
	}
}