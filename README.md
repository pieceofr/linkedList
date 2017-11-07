This is single linkedlist implement by GO.

You can store any DataType in LinkList as long as you make an struct which implemnt interface

Follows is how you doing it:

In the **listype.go**, you may define your struct like this

```
type MyType struct {
	myvalue int
	mystruct Node
}
```

You need to implement **equal method** to fullfill **ListVal interface**

```
func (m MyType)equal(val ListVal) bool {
	if m.myvalue == val.(MyType).myvalue {
		return true
	}
	return false
}
```

Finally, doing your testing on **linkedlist_test.go** and check **func TestListVal(t *testing.T)**, you can do the following test

```
testMyType1 := MyType{myvalue:1, mystruct:Node{value:INTVAL(5)}}
testMyType2 := MyType{myvalue:1, mystruct:Node{value:INTVAL(5)}}
testMyType3 := MyType{myvalue:2, mystruct:Node{value:INTVAL(5)}}

if !testMyType1.equal(testMyType2) {
    t.Error("Should be eqaul!")
}
if testMyType1.equal(testMyType3) {
    t.Error("Should be not eqaul!")
}
````

-----------------------


## Linked List Operations (1)
+ AddFirst
![](https://i.imgur.com/0oqFQEB.jpg)

+ Travesering
![](https://i.imgur.com/jl1hDvn.jpg)

----

## Linked List Operations (2)

+ AddLast
![](https://i.imgur.com/szOdYyY.jpg)

+ Inserting "after"
![](https://i.imgur.com/IvWdSZi.jpg)

----

## Linked List Operations (3)

+ Inserting "before"
![](https://i.imgur.com/1ZHlWb1.jpg)

+ Deletion
![](https://i.imgur.com/D17PUdk.jpg)

---

## Iterator / Cursor

The whole idea of the iterator is to provide an access to a private aggregated data and at the same moment hiding the underlying representation

----

## Functions of Iterator
+ AnyType next() - returns the next element in the container
+ boolean hasNext() - checks if there is a next element
+ void remove() - (optional operation).removes the element returned by next()

---

# The Runner Technique
+ Detecting a circular linked list
+ One fast and one slow, eventually they will be in the same node sometimes

---

## Reference

+ [www.cs.cmu.edu Linked Lists](https://www.cs.cmu.edu/~adamchik/15-121/lectures/Linked%20Lists/linked%20lists.html)

+ [Linked Lists By Gayle Laakmann ](https://www.youtube.com/watch?v=njTh_OwMljA)

+ [Cycles in a Linked List](https://www.youtube.com/watch?v=MFOAbpfrJ8g)

+ [My Go Implement LinkedList](https://github.com/pieceofr/linkedList)

###### tags: `wwc` `cracking the code` `linked lists`