package rbtree

import (
	"errors"
	"fmt"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/visitor"
)

var ErrorNotFound = errors.New("not found")

// RbTree is a kind of self-balancing binary search tree in computer science.
// Each node of the binary tree has an extra bit, and that bit is often interpreted
// as the color (red or black) of the node. These color bits are used to ensure the tree
// remains approximately balanced during insertions and deletions.
type RbTree[K, V any] struct {
	root   *Node[K, V]
	size   int
	keyCmp comparator.Comparator[K]
}

// New creates a new RbTree
func New[K, V any](cmp comparator.Comparator[K]) *RbTree[K, V] {
	return &RbTree[K, V]{keyCmp: cmp}
}

// Clear clears the RbTree
func (t *RbTree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

// Find finds the first node that the key is equal to the passed key, and returns its value
func (t *RbTree[K, V]) Find(key K) (V, error) {
	n := t.findFirstNode(key)
	if n != nil {
		return n.value, nil
	}
	return *new(V), ErrorNotFound
}

// FindNode the first node that the key is equal to the passed key and return it
func (t *RbTree[K, V]) FindNode(key K) *Node[K, V] {
	return t.findFirstNode(key)
}

// Begin returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) Begin() *Node[K, V] {
	return t.First()
}

// First returns the node with minimum key in the RbTree
func (t *RbTree[K, V]) First() *Node[K, V] {
	if t.root == nil {
		return nil
	}
	return minimum(t.root)
}

// RBegin returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) RBegin() *Node[K, V] {
	return t.Last()
}

// Last returns the Node with maximum key in the RbTree
func (t *RbTree[K, V]) Last() *Node[K, V] {
	if t.root == nil {
		return nil
	}
	return maximum(t.root)
}

// IterFirst returns the iterator of first node
func (t *RbTree[K, V]) IterFirst() *RbTreeIterator[K, V] {
	return NewIterator(t.First())
}

// IterLast returns the iterator of first node
func (t *RbTree[K, V]) IterLast() *RbTreeIterator[K, V] {
	return NewIterator(t.Last())
}

// Empty returns true if Tree is empty,otherwise returns false.
func (t *RbTree[K, V]) Empty() bool {
	return t.size == 0
}

// Size returns the size of the rbtree.
func (t *RbTree[K, V]) Size() int {
	return t.size
}

// Insert inserts a key-value pair into the RbTree.
func (t *RbTree[K, V]) Insert(key K, value V) {
	x := t.root
	var y *Node[K, V]

	for x != nil {
		y = x
		if t.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node[K, V]{parent: y, color: RED, key: key, value: value}
	t.size++

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if t.keyCmp(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)
}

func (t *RbTree[K, V]) rbInsertFixup(z *Node[K, V]) {
	var y *Node[K, V]
	for z.parent != nil && !z.parent.color {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && !y.color {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && !y.color {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// Delete deletes node from the RbTree
func (t *RbTree[K, V]) Delete(node *Node[K, V]) {
	z := node
	if z == nil {
		return
	}

	var x, y *Node[K, V]
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
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

func (t *RbTree[K, V]) rbDeleteFixup(x, parent *Node[K, V]) {
	var w *Node[K, V]
	for x != t.root && getColor(x) {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			x, w = t.rbFixupLeft(x, parent, w)
		} else {
			x, w = t.rbFixupRight(x, parent, w)
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *RbTree[K, V]) rbFixupLeft(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.right
	if !w.color {
		w.color = BLACK
		parent.color = RED
		t.leftRotate(parent)
		w = parent.right
	}
	if getColor(w.left) && getColor(w.right) {
		w.color = RED
		x = parent
	} else {
		if getColor(w.right) {
			if w.left != nil {
				w.left.color = BLACK
			}
			w.color = RED
			t.rightRotate(w)
			w = parent.right
		}
		w.color = parent.color
		parent.color = BLACK
		if w.right != nil {
			w.right.color = BLACK
		}
		t.leftRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree[K, V]) rbFixupRight(x, parent, w *Node[K, V]) (*Node[K, V], *Node[K, V]) {
	w = parent.left
	if !w.color {
		w.color = BLACK
		parent.color = RED
		t.rightRotate(parent)
		w = parent.left
	}
	if getColor(w.left) && getColor(w.right) {
		w.color = RED
		x = parent
	} else {
		if getColor(w.left) {
			if w.right != nil {
				w.right.color = BLACK
			}
			w.color = RED
			t.leftRotate(w)
			w = parent.left
		}
		w.color = parent.color
		parent.color = BLACK
		if w.left != nil {
			w.left.color = BLACK
		}
		t.rightRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree[K, V]) leftRotate(x *Node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RbTree[K, V]) rightRotate(x *Node[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// findNode finds the node that its key is equal to the passed key, and returns it.
func (t *RbTree[K, V]) findNode(key K) *Node[K, V] {
	x := t.root
	for x != nil {
		if t.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			if t.keyCmp(key, x.key) == 0 {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// findNode finds the first node that its key is equal to the passed key, and returns it
func (t *RbTree[K, V]) findFirstNode(key K) *Node[K, V] {
	node := t.FindLowerBoundNode(key)
	if node == nil {
		return nil
	}
	if t.keyCmp(node.key, key) == 0 {
		return node
	}
	return nil
}

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (t *RbTree[K, V]) FindLowerBoundNode(key K) *Node[K, V] {
	return t.findLowerBoundNode(t.root, key)
}

func (t *RbTree[K, V]) findLowerBoundNode(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	if t.keyCmp(key, x.key) <= 0 {
		ret := t.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		}
		if t.keyCmp(ret.key, x.key) <= 0 {
			return ret
		}
		return x
	}
	return t.findLowerBoundNode(x.right, key)
}

// FindUpperBoundNode finds the first node that its key is greater than the passed key, and returns it
func (t *RbTree[K, V]) FindUpperBoundNode(key K) *Node[K, V] {
	return t.findUpperBoundNode(t.root, key)
}

func (t *RbTree[K, V]) findUpperBoundNode(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	if t.keyCmp(key, x.key) >= 0 {
		return t.findUpperBoundNode(x.right, key)
	}
	ret := t.findUpperBoundNode(x.left, key)
	if ret == nil {
		return x
	}
	if t.keyCmp(ret.key, x.key) <= 0 {
		return ret
	}
	return x
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (t *RbTree[K, V]) Traversal(visitor visitor.KvVisitor[K, V]) {
	for node := t.First(); node != nil; node = node.Next() {
		if !visitor(node.key, node.value) {
			break
		}
	}
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (t *RbTree[K, V]) IsRbTree() (bool, error) {
	// Properties:
	// 1. Each node is either red or black.
	// 2. The root is black.
	// 3. All leaves (NIL) are black.
	// 4. If a node is red, then both its children are black.
	// 5. Every path from a given node to any of its descendant NIL nodes contains the same number of black nodes.
	_, property, ok := t.test(t.root)
	if !ok {
		return false, fmt.Errorf("violate property %v", property)
	}
	return true, nil
}

func (t *RbTree[K, V]) test(n *Node[K, V]) (int, int, bool) {

	if n == nil { // property 3:
		return 1, 0, true
	}

	if n == t.root && !n.color { // property 2:
		return 1, 2, false
	}
	leftBlackCount, property, ok := t.test(n.left)
	if !ok {
		return leftBlackCount, property, ok
	}
	rightBlackCount, property, ok := t.test(n.right)
	if !ok {
		return rightBlackCount, property, ok
	}

	if rightBlackCount != leftBlackCount { // property 5:
		return leftBlackCount, 5, false
	}
	blackCount := leftBlackCount

	if !n.color {
		if !getColor(n.left) || !getColor(n.right) { // property 4:
			return 0, 4, false
		}
	} else {
		blackCount++
	}

	// if n == t.root {
	// 	fmt.Printf("blackCount:%v \n", blackCount)
	// }
	return blackCount, 0, true
}

// getColor returns the node's color
func getColor[K, V any](n *Node[K, V]) Color {
	if n == nil {
		return BLACK
	}
	return n.color
}
