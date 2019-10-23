package rbtree

import (
	. "github.com/liyue201/gostl/comparator"
	. "github.com/liyue201/gostl/iterator"
) 

type RbTree struct {
	root *Node
	size int
	cmp  Comparator
}

func New(cmp Comparator) *RbTree {
	return &RbTree{cmp: cmp}
}

// Clear clears the tree
func (this *RbTree) Clear() {
	this.root = nil
	this.size = 0
}

// Find finds the first Node by the key and return its value.
func (this *RbTree) Find(key interface{}) interface{} {
	n := this.findFirstNode(key)
	if n != nil {
		return n.value
	}
	return nil
}

// FindIt finds the first Node and return it as an iterator.
func (this *RbTree) FindNode(key interface{}) *Node {
	return this.findFirstNode(key)
}

// Begin returns the Node with minimum key in the tree
func (this *RbTree) Begin() *Node {
	return this.First()
}

// Fisrt returns the Node with minimum key in the tree
func (this *RbTree) First() *Node {
	if this.root == nil {
		return nil
	}
	return minimum(this.root)
}

// RBegin returns the Node with maximum key in the tree
func (this *RbTree) RBegin() *Node {
	return this.Last()
}

// Last returns the Node with maximum key in the tree
func (this *RbTree) Last() *Node {
	if this.root == nil {
		return nil
	}
	return maximum(this.root)
}

// IterFirst returns the iterator of first Node
func (this *RbTree) IterFirst() KvBidIterator {
	return NewIterator(this.First())
}

// IterLast returns the iterator of first Node
func (this *RbTree) IterLast() KvBidIterator {
	return NewIterator(this.Last())
}

// Empty returns true if Tree is empty,otherwise returns false.
func (this *RbTree) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

// Size returns the size of the rbtree.
func (this *RbTree) Size() int {
	return this.size
}

// Insert inserts a key-value pair into the rbtree.
func (this *RbTree) Insert(key, value interface{}) {
	x := this.root
	var y *Node

	for x != nil {
		y = x
		if this.cmp(key, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, color: RED, key: key, value: value}
	this.size++

	if y == nil {
		z.color = BLACK
		this.root = z
		return
	} else if this.cmp(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	this.rbInsertFixup(z)
}

func (this *RbTree) rbInsertFixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					this.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				this.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					this.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				this.leftRotate(z.parent.parent)
			}
		}
	}
	this.root.color = BLACK
}

// Delete deletes the Node
func (this *RbTree) Delete(node *Node) {
	z := node
	if z == nil {
		return
	}

	var x, y *Node
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		this.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color == BLACK {
		this.rbDeleteFixup(x, xparent)
	}
	this.size--
}

func (this *RbTree) rbDeleteFixup(x, parent *Node) {
	var w *Node

	for x != this.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				this.leftRotate(parent)
				w = parent.right
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.right) == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					this.rightRotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				this.leftRotate(parent)
				x = this.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				this.rightRotate(parent)
				w = parent.left
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.left) == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					this.leftRotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				this.rightRotate(parent)
				x = this.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (this *RbTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		this.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (this *RbTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		this.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// findNode finds the Node by key and return it's Node, if not exists return nil.
func (this *RbTree) findNode(key interface{}) *Node {
	x := this.root
	for x != nil {
		if this.cmp(key, x.key) < 0 {
			x = x.left
		} else {
			if this.cmp(key, x.key) == 0 {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// findNode returns the first Node that equal to key, if not exists return nil.
func (this *RbTree) findFirstNode(key interface{}) *Node {
	node := this.FindLowerBoundNode(key)
	if node == nil {
		return nil
	}
	if this.cmp(node.key, key) == 0 {
		return node
	}
	return nil
}

// findNode returns the first Node that equal or greater than key, if not exists return nil.
func (this *RbTree) FindLowerBoundNode(key interface{}) *Node {
	return this.findLowerBoundNode(this.root, key)
}

func (this *RbTree) findLowerBoundNode(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	if this.cmp(key, x.key) <= 0 {
		ret := this.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		} else {
			if this.cmp(ret.key, x.key) <= 0 {
				return ret
			} else {
				return x
			}
		}
	} else {
		return this.findLowerBoundNode(x.right, key)
	}
}

// getColor gets color of the Node.
func getColor(n *Node) Color {
	if n == nil {
		return BLACK
	}
	return n.color
}
