package linkedlist

/*I do'nt have testing for this one yet*/
type ListIterator struct {
	nextNode *Node	
}

func  (i *ListIterator)GetElement() *Node {
	return i.nextNode
}

func  (i *ListIterator)Next() *Node {
	if i == nil {
		return nil
	}
	i.nextNode = i.nextNode.next
	return i.nextNode
}

func  (i *ListIterator)hasNext() bool {
	if i.nextNode.next == nil	{
		return false
	} else {
		return true
	}
}