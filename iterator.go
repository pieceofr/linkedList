package linkedlist

import "errors"

/*I do'nt have testing for this one yet*/
type ListIterator struct {
	nextNode *Node	
}

func (i *ListIterator)New(list LinkedList) (ListIterator) {
	return ListIterator{nextNode:list.head}
}

func  (i *ListIterator)Next() *Node {
	if i == nil {
		return nil
	}
	i.nextNode = i.nextNode.next
	return i.nextNode
}

func  (i *ListIterator)hasNext() (bool, error) {
	if i == nil {
		return false, errors.New("ListIterator is not initialized")
	}
	if i.nextNode.next == nil	{
		return false, nil
	} else {
		return true, nil
	}
}