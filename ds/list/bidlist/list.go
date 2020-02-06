package bidlist

import (
	"fmt"
	"github.com/liyue201/gostl/ds/container"
	"github.com/liyue201/gostl/utils/visitor"
)

// List is an implementation of Container
var _ container.Container = (*List)(nil)

// Node is a list node
type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
	list  *List
}

// Next returns the next list node or nil.
func (n *Node) Next() *Node {
	if n.list == nil {
		return nil
	}
	if n.next == n.list.head {
		return nil
	}
	return n.next
}

// Prev returns the previous list node or nil.
func (n *Node) Prev() *Node {
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
type List struct {
	head *Node // point to the front Node
	len  int   // current list length
}

// New creates a list
func New() *List {
	list := &List{}
	return list
}

// Len returns the amount of list nodes.
func (l *List) Len() int {
	return l.len
}

// Size returns the amount of list nodes.
func (l *List) Size() int {
	return l.len
}

// Empty returns true if the list is empty
func (l *List) Empty() bool {
	return l.len == 0
}

// FrontNode returns the front node of the list or nil if the list is empty
func (l *List) FrontNode() *Node {
	return l.head
}

// BackNode returns the last node of the list or nil if the list is empty
func (l *List) BackNode() *Node {
	if l.head == nil {
		return nil
	}
	return l.head.prev
}

// Front returns the value of the front node
func (l *List) Front() interface{} {
	if l.len == 0 {
		return nil
	}
	return l.head.Value
}

// Back returns the value of the last node
func (l *List) Back() interface{} {
	if l.len == 0 {
		return nil
	}
	return l.head.prev.Value
}

// PushBack inserts a new node n with value v at the back of the list
func (l *List) PushBack(v interface{}) {
	l.pushBack(v)
}

// PushBack inserts a new node n with value v at the back of the list and returns n.
func (l *List) pushBack(v interface{}) *Node {
	n := &Node{Value: v, list: l}
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
func (l *List) PushFront(v interface{}) {
	n := l.pushBack(v)
	l.head = n
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertAfter(&Node{Value: v, list: l}, mark)
}

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of l list, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	n := l.insertAfter(&Node{Value: v, list: l}, mark.prev)
	if l.head == mark {
		l.head = n
	}
	return n
}

func (l *List) insertAfter(n, at *Node) *Node {
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
func (l *List) Remove(n *Node) interface{} {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}

func (l *List) remove(n *Node) *Node {
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
func (l *List) Clear() {
	l.head = nil
	l.len = 0
}

// PopBack removes the last node in the list and returns its value
func (l *List) PopBack() interface{} {
	n := l.BackNode()
	if n != nil {
		return l.Remove(n)
	}
	return nil
}

// PopFront removes the first node in the list and returns its value
func (l *List) PopFront() interface{} {
	n := l.FrontNode()
	if n != nil {
		return l.Remove(n)
	}
	return nil
}

// MoveToFront moves node n to the front of the list.
// If n is not a node of the list, the list is not modified.
// The n must not be nil.
func (l *List) MoveToFront(n *Node) {
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
func (l *List) MoveToBack(n *Node) {
	if n.list != l {
		return
	}
	if l.head.prev != n {
		l.moveToAfter(n, l.head.prev)
	}
}

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of the list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (l *List) MoveAfter(n, mark *Node) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.moveToAfter(n, mark)
}

func (l *List) moveToAfter(n, at *Node) {
	if n == at.next {
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
func (l *List) PushBackList(other *List) {
	for i, n := other.Len(), other.FrontNode(); i > 0; i, n = i-1, n.Next() {
		l.InsertAfter(n.Value, l.head.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of the list.
// The list and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	for i, e := other.Len(), other.BackNode(); i > 0; i, e = i-1, e.Prev() {
		l.InsertBefore(e.Value, l.head)
	}
}

// String returns a string representation of the list
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

// Traversal traversals elements in the list, it will not stop until to the end of the list or the visitor returns false
func (l *List) Traversal(visitor visitor.Visitor) {
	for node := l.FrontNode(); node != nil; node = node.Next() {
		if !visitor(node.Value) {
			break
		}
	}
}
