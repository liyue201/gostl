package simplelist

import (
	"fmt"
	"github.com/liyue201/gostl/utils/visitor"
)

// Node is a list node
type Node struct {
	next  *Node
	Value interface{}
}

// Next returns the next list node or nil.
func (n *Node) Next() *Node {
	return n.next
}

// List represents a single direction list:
//
//   head -> node1 --> node2 --> node3 <- tail
//
type List struct {
	head *Node // point to the front Node
	tail *Node // point to the back Node
	len  int   // current list length
}

// New news a list
func New() *List {
	list := &List{}
	return list
}

// Len returns the number of nodes of list.
func (l *List) Len() int {
	return l.len
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List) FrontNode() *Node {
	return l.head
}

// BackNode returns the lase node of the list or nil if the list is empty
func (l *List) BackNode() *Node {
	return l.tail
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List) PushFront(v interface{}) {
	n := &Node{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.len++
}

// PushBack inserts a new node n with value v at the back of the list.
func (l *List) PushBack(v interface{}) {
	n := &Node{Value: v}
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len++
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	return l.insertAfter(&Node{Value: v}, mark)
}

func (l *List) insertAfter(n, at *Node) *Node {
	n.next = at.next
	at.next = n
	if n.next == nil {
		l.tail = n
	}
	l.len++
	return n
}

// Remove removes node n from l list.
// The node must not be nil.
func (l *List) Remove(pre, n *Node) interface{} {
	if n == nil {
		return nil
	}
	if pre == nil {
		l.head = n.next
		if l.head == nil {
			l.tail = nil
		}
	} else {
		pre.next = n.next
		if pre.next == nil {
			l.tail = pre
		}
	}
	l.len--
	return n.Value
}

// MoveToFront moves node n to the front of l list.
// The n must not be nil.
func (l *List) MoveToFront(pre, n *Node) {
	if pre == nil || pre.next != n || n == nil || l.len <= 1 {
		return
	}
	pre.next = n.next
	if pre.next == nil {
		l.tail = pre
	}
	n.next = l.head
	l.head = n
}

// MoveToBack moves node n to the back of l list.
// The n must not be nil.
func (l *List) MoveToBack(pre, n *Node) {
	if n == nil || n.next == nil || l.len <= 1 {
		return
	}
	if pre == nil {
		l.head = n.next
	} else {
		pre.next = n.next
	}
	l.tail.next = n
	l.tail = n
}

// String returns the list content in string format
func (l *List) String() string {
	str := "["
	for n := l.FrontNode(); n != nil; n = n.Next() {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", n.Value)
	}
	str += "]"
	return str
}

// Traversal traversals elements in list, it will not stop until to the end or visitor returns false
func (l *List) Traversal(visitor visitor.Visitor) {
	for node := l.head; node != nil; node = node.next {
		if !visitor(node.Value) {
			break
		}
	}
}
