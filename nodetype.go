package linkedlist

type INTVAL int
type STRVAL string

type NodeType interface {
	equal(NodeType) bool
}

func (i INTVAL)equal(val NodeType) bool {
	if int(i) == int(val.(INTVAL)) {
		return true
	}
	return false
}

func (s STRVAL)equal(val NodeType) bool {
	if string(s) == string(val.(STRVAL)) {
		return true
	}
	return false
}

type MyType struct {
	myvalue int
	mystruct Node
}

func (m MyType)equal(val NodeType) bool {
	if m.myvalue == val.(MyType).myvalue {
		return true
	}
	return false
}