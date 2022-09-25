package simplelist

import (
	"fmt"
	"github.com/liyue201/gostl/utils/visitor"
)

// Node is a list node
type Node[T any] struct {
	next  *Node[T]
	Value T
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// List represents a single direction list:
//
//   head -> node1 --> node2 --> node3 <- tail
//
type List[T any] struct {
	head *Node[T] // point to the front Node
	tail *Node[T] // point to the back Node
	len  int      // current list length
}

// New creates a list
func New[T any]() *List[T] {
	list := &List[T]{}
	return list
}

// Len returns the amount of list nodes.
func (l *List[T]) Len() int {
	return l.len
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List[T]) FrontNode() *Node[T] {
	return l.head
}

// BackNode returns the last node of the list or nil if the list is empty
func (l *List[T]) BackNode() *Node[T] {
	return l.tail
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List[T]) PushFront(v T) {
	n := &Node[T]{Value: v}
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
func (l *List[T]) PushBack(v T) {
	n := &Node[T]{Value: v}
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
// If mark is not a node of the list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] {
	return l.insertAfter(&Node[T]{Value: v}, mark)
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	n.next = at.next
	at.next = n
	if n.next == nil {
		l.tail = n
	}
	l.len++
	return n
}

// Remove removes node n from the list.
// The node must not be nil.
func (l *List[T]) Remove(pre, n *Node[T]) T {
	if n == nil {
		return *new(T)
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

// MoveToFront moves node n to the front of the list.
// The n must not be nil.
func (l *List[T]) MoveToFront(pre, n *Node[T]) {
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

// MoveToBack moves node n to the back of the list.
// The n must not be nil.
func (l *List[T]) MoveToBack(pre, n *Node[T]) {
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
	n.next = nil
}

// String returns a string representation of the list
func (l *List[T]) String() string {
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

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List[T]) Traversal(visitor visitor.Visitor[T]) {
	for node := l.head; node != nil; node = node.Next() {
		if !visitor(node.Value) {
			break
		}
	}
}
