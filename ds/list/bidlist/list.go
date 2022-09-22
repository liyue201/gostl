package bidlist

import (
	"errors"
	"fmt"
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/utils/visitor"
)

// List is an implementation of Container
type T any

var ErrorOutOffRange = errors.New("out of range")

var _ container.Container[T] = (*List[T])(nil)

// Node is a list node
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	Value T
	list  *List[T]
}

// Next returns the next list node or nil.
func (n *Node[T]) Next() *Node[T] {
	if n.list == nil {
		return nil
	}
	if n.next == n.list.head {
		return nil
	}
	return n.next
}

// Prev returns the previous list node or nil.
func (n *Node[T]) Prev() *Node[T] {
	if n.list == nil {
		return nil
	}
	if n == n.list.head {
		return nil
	}
	return n.prev
}

// List represents a bidirectional list:
//
//   head -> node1 -- node2 --  node3
//            |                   |
//           node6 -- node5 --  node4
//
type List[T any] struct {
	head *Node[T] // point to the front Node
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

// Size returns the amount of list nodes.
func (l *List[T]) Size() int {
	return l.len
}

// Empty returns true if the list is empty
func (l *List[T]) Empty() bool {
	return l.len == 0
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List[T]) FrontNode() *Node[T] {
	return l.head
}

// BackNode returns the last node of the list or nil if the list is empty
func (l *List[T]) BackNode() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.head.prev
}

// Front returns the value of the front node
func (l *List[T]) Front() T {
	if l.len == 0 {
		panic(ErrorOutOffRange)
	}
	return l.head.Value
}

// Back returns the value of the last node
func (l *List[T]) Back() T {
	if l.len == 0 {
		panic(ErrorOutOffRange)
	}
	return l.head.prev.Value
}

// PushBack inserts a new node n with value v at the back of the list
func (l *List[T]) PushBack(v T) {
	l.pushBack(v)
}

// PushBack inserts a new node n with value v at the back of the list and returns n.
func (l *List[T]) pushBack(v T) *Node[T] {
	n := &Node[T]{Value: v, list: l}
	if l.len == 0 {
		n.prev = n
		n.next = n
		l.head = n
		l.len++
		return n
	}
	return l.insertAfter(n, l.head.prev)
}

// PushFront inserts a new node n with value v at the front of the list.
func (l *List[T]) PushFront(v T) {
	n := l.pushBack(v)
	l.head = n
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Node[T]) *Node[T] {
	if mark.list != l {
		return nil
	}
	return l.insertAfter(&Node[T]{Value: v, list: l}, mark)
}

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Node[T]) *Node[T] {
	if mark.list != l {
		return nil
	}
	n := l.insertAfter(&Node[T]{Value: v, list: l}, mark.prev)
	if l.head == mark {
		l.head = n
	}
	return n
}

func (l *List[T]) insertAfter(n, at *Node[T]) *Node[T] {
	n.next = at.next
	n.prev = at
	at.next.prev = n
	at.next = n
	l.len++
	return n
}

// Remove removes n from l list if n is a node of l list.
// It returns the n value n.Value.
// The node must not be nil.
func (l *List[T]) Remove(n *Node[T]) T {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}

func (l *List[T]) remove(n *Node[T]) *Node[T] {
	if n == l.head {
		l.head = l.head.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil // avoid memory leaks
	n.prev = nil // avoid memory leaks
	n.list = nil
	l.len--
	if l.len == 0 {
		l.head = nil
	}
	return n
}

// Clear removes all nodes
func (l *List[T]) Clear() {
	l.head = nil
	l.len = 0
}

// PopBack removes the last node in the list and returns its value
func (l *List[T]) PopBack() T {
	n := l.BackNode()
	if n != nil {
		return l.Remove(n)
	}
	panic(ErrorOutOffRange)
}

// PopFront removes the first node in the list and returns its value
func (l *List[T]) PopFront() T {
	n := l.FrontNode()
	if n != nil {
		return l.Remove(n)
	}
	panic(ErrorOutOffRange)
}

// MoveToFront moves node n to the front of the list.
// If n is not a node of the list, the list is not modified.
// The n must not be nil.
func (l *List[T]) MoveToFront(n *Node[T]) {
	if n.list != l {
		return
	}
	if l.head != n {
		l.moveToAfter(n, l.head.prev)
		l.head = n
	}
}

// MoveToBack moves node  n to the back of the list.
// If e is not a node of the list, the list is not modified.
// The node must not be nil.
func (l *List[T]) MoveToBack(n *Node[T]) {
	if n.list != l {
		return
	}
	if l.head.prev != n {
		if l.head == n {
			l.head = n.next
		}
		l.moveToAfter(n, l.head.prev)
	}
}

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of the list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List[T]) MoveAfter(n, mark *Node[T]) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.moveToAfter(n, mark)
}

func (l *List[T]) moveToAfter(n, at *Node[T]) {
	if n == at.next || n == at {
		return
	}
	if n == l.head {
		l.head = n.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = at.next
	n.prev = at
	at.next.prev = n
	at.next = n
}

// PushBackList inserts a copy of an other list at the back of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) {
	for i, n := other.Len(), other.FrontNode(); i > 0; i, n = i-1, n.Next() {
		l.InsertAfter(n.Value, l.head.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of the list.
// The list and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
	for i, e := other.Len(), other.BackNode(); i > 0; i, e = i-1, e.Prev() {
		l.InsertBefore(e.Value, l.head)
	}
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
func (l *List[T]) Traversal(visitor visitor.Visitor) {
	for node := l.FrontNode(); node != nil; node = node.Next() {
		if !visitor(node.Value) {
			break
		}
	}
}
