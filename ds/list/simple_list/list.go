package simple_list

import (
	"fmt"
	"github.com/liyue201/gostl/utils/visitor"
)

type Node struct {
	next  *Node
	Value interface{}
}

// Next returns the next list node or nil.
func (this *Node) Next() *Node {
	return this.next
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

func New() *List {
	list := &List{}
	return list
}

// Len returns the number of nodes of list.
func (this *List) Len() int {
	return this.len
}

// FrontNode returns the front node of the list or nil if the list is empty
func (this *List) FrontNode() *Node {
	return this.head
}

// BackNode returns the lase node of the list or nil if the list is empty
func (this *List) BackNode() *Node {
	return this.tail
}

// PushFront inserts a new node n with value v at the front of the list.
func (this *List) PushFront(v interface{}) {
	n := &Node{Value: v}
	if this.len == 0 {
		this.head = n
		this.tail = n
	} else {
		n.next = this.head
		this.head = n
	}
	this.len++
}

// PushBack inserts a new node n with value v at the back of the list.
func (this *List) PushBack(v interface{}) {
	n := &Node{Value: v}
	if this.len == 0 {
		this.head = n
		this.tail = n
	} else {
		this.tail.next = n
		this.tail = n
	}
	this.len++
}

// InsertAfter inserts a new node n with value v immediately after mark and returns n.
// If mark is not a node of this list, the list is not modified.
// The mark must not be nil.
func (this *List) InsertAfter(v interface{}, mark *Node) *Node {
	return this.insertAfter(&Node{Value: v}, mark)
}

func (this *List) insertAfter(n, at *Node) *Node {
	n.next = at.next
	at.next = n
	if n.next == nil {
		this.tail = n
	}
	this.len++
	return n
}

// Remove removes node n from this list.
// The node must not be nil.
func (this *List) Remove(pre, n *Node) interface{} {
	if n == nil {
		return nil
	}
	if pre == nil {
		this.head = n.next
		if this.head == nil {
			this.tail = nil
		}
	} else {
		pre.next = n.next
		if pre.next == nil {
			this.tail = pre
		}
	}
	this.len--
	return n.Value
}

// MoveToFront moves node n to the front of this list.
// The n must not be nil.
func (this *List) MoveToFront(pre, n *Node) {
	if pre == nil || pre.next != n || n == nil || this.len <= 1 {
		return
	}
	pre.next = n.next
	if pre.next == nil {
		this.tail = pre
	}
	n.next = this.head
	this.head = n
}

// MoveToBack moves node n to the back of this list.
// The n must not be nil.
func (this *List) MoveToBack(pre, n *Node) {
	if n == nil || n.next == nil || this.len <= 1 {
		return
	}
	if pre == nil {
		this.head = n.next
	} else {
		pre.next = n.next
	}
	this.tail.next = n
	this.tail = n
}

// String returns the list content in string format
func (this *List) String() string {
	str := "["
	for n := this.FrontNode(); n != nil; n = n.Next() {
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
