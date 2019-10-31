package bid_list

import (
	"fmt"
	"github.com/liyue201/gostl/utils/visitor"
)

type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
	list  *List
}

// Next returns the next list node or nil.
func (this *Node) Next() *Node {
	if this.list == nil {
		return nil
	}
	if this.next == this.list.head {
		return nil
	}
	return this.next
}

// Prev returns the previous list node or nil.
func (this *Node) Prev() *Node {
	if this.list == nil {
		return nil
	}
	if this == this.list.head {
		return nil
	}
	return this.prev
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

func New() *List {
	list := &List{}
	return list
}

// Len returns the number of nodes of list.
func (this *List) Len() int {
	return this.len
}

// Front returns the front node of the list or nil if the list is empty
func (this *List) Front() *Node {
	if this.len == 0 {
		return nil
	}
	return this.head
}

// Front returns the lase node of the list or nil if the list is empty
func (this *List) Back() *Node {
	if this.len == 0 {
		return nil
	}
	return this.head.prev
}

// PushBack inserts a new node n with value v at the back of the list and returns n.
func (this *List) PushBack(v interface{}) *Node {
	n := &Node{Value: v, list: this}
	if this.len == 0 {
		n.prev = n
		n.next = n
		this.head = n
		this.len++
		return n
	}
	return this.insertAfter(n, this.head.prev)
}

// PushFront inserts a new node n with value v at the front of the list and returns n.
func (this *List) PushFront(v interface{}) *Node {
	n := this.PushBack(v)
	this.head = n
	return n
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of this list, the list is not modified.
// The mark must not be nil.
func (this *List) InsertAfter(v interface{}, mark *Node) *Node {
	if mark.list != this {
		return nil
	}
	return this.insertAfter(&Node{Value: v, list: this}, mark)
}

// InsertBefore inserts a new node n with value v immediately before mark and returns n.
// If mark is not a node of this list, the list is not modified.
// The mark must not be nil.
func (this *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark.list != this {
		return nil
	}
	n := this.insertAfter(&Node{Value: v, list: this}, mark.prev)
	if this.head == mark {
		this.head = n
	}
	return n
}

func (this *List) insertAfter(n, at *Node) *Node {
	n.next = at.next
	n.prev = at
	at.next.prev = n
	at.next = n
	this.len++
	return n
}

// Remove removes n from this list if n is a node of this list.
// It returns the n value n.Value.
// The node must not be nil.
func (this *List) Remove(n *Node) interface{} {
	if n.list == this {
		this.remove(n)
	}
	return n.Value
}

func (this *List) remove(n *Node) *Node {
	if n == this.head {
		this.head = this.head.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil // avoid memory leaks
	n.prev = nil // avoid memory leaks
	n.list = nil
	this.len--
	if this.len == 0 {
		this.head = nil
	}
	return n
}

// MoveToFront moves node n to the front of this list.
// If n is not a node of this list, the list is not modified.
// The n must not be nil.
func (this *List) MoveToFront(n *Node) {
	if n.list != this {
		return
	}
	if this.head != n {
		this.moveToAfter(n, this.head.prev)
		this.head = n
	}
}

// MoveToBack moves node  n to the back of this list.
// If e is not a node of this list, the list is not modified.
// The node must not be nil.
func (this *List) MoveToBack(n *Node) {
	if n.list != this {
		return
	}
	if this.head.prev != n {
		this.moveToAfter(n, this.head.prev)
	}
}

// MoveAfter moves node n to its new position after mark.
// If n or mark is not a node of this list, or n == mark, the list is not modified.
// The node and mark must not be nil.
func (this *List) MoveAfter(n, mark *Node) {
	if n.list != this || n == mark || mark.list != this {
		return
	}
	this.moveToAfter(n, mark)
}

func (this *List) moveToAfter(n, at *Node) {
	if n == at.next {
		return
	}
	if n == this.head {
		this.head = n.next
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = at.next
	n.prev = at
	at.next.prev = n
	at.next = n
}

// PushBackList inserts a copy of an other list at the back of this list.
// The this list and other may be the same. They must not be nil.
func (this *List) PushBackList(other *List) {
	for i, n := other.Len(), other.Front(); i > 0; i, n = i-1, n.Next() {
		this.InsertAfter(n.Value, this.head.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of this list.
// The this list and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.InsertBefore(e.Value, l.head)
	}
}

// String returns the list content in string format
func (this *List) String() string {
	str := "["
	for n := this.Front(); n != nil; n = n.Next() {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", n.Value)
	}
	str += "]"
	return str
}

// Traversal traversals elements in list, it will not stop until to the end or visitor returns false
func (this *List) Traversal(visitor visitor.Visitor) {
	for node := this.head; node != nil; node = node.next {
		if !visitor(node.Value) {
			break
		}
	}
}
