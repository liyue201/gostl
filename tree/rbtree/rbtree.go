package rbtree

const (
	RED   = 0
	BLACK = 1
)

// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator func(a, b interface{}) int

type node struct {
	left, right, parent *node
	color               int
	Key                 interface{}
	Value               interface{}
}

type RbTree struct {
	root *node
	size int
	cmp  Comparator
}

func New(cmp Comparator) *RbTree {
	return &RbTree{cmp: cmp}
}

// Find finds the node bye the key and return its value.
func (t *RbTree) Find(key interface{}) interface{} {
	n := t.findNode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

// FindIt finds the node and return it as an iterator.
//func (t *RbTree) FindIt(key interface{}) *node {
//	return t.findNode(key)
//}

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
	var y *node

	for x != nil {
		y = x
		if this.cmp(key, x.Key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &node{parent: y, color: RED, Key: key, Value: value}
	this.size++

	if y == nil {
		z.color = BLACK
		this.root = z
		return
	} else if this.cmp(z.Key, y.Key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	this.rbInsertFixup(z)
}

func (this *RbTree) rbInsertFixup(z *node) {
	var y *node
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

// Delete deletes the node by key
func (this *RbTree) Delete(key interface{}) {
	z := this.findNode(key)
	if z == nil {
		return
	}

	var x, y *node
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
		z.Key = y.Key
		z.Value = y.Value
	}

	if y.color == BLACK {
		this.rbDeleteFixup(x, xparent)
	}
	this.size--
}

func (this *RbTree) rbDeleteFixup(x, parent *node) {
	var w *node

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

func (this *RbTree) leftRotate(x *node) {
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

func (this *RbTree) rightRotate(x *node) {
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

// findNode finds the node by key and return it, if not exists return nil.
func (this *RbTree) findNode(key interface{}) *node {
	x := this.root
	for x != nil {
		if this.cmp(key, x.Key) < 0 {
			x = x.left
		} else {
			if this.cmp(key, x.Key) == 0 {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// Next returns the node's successor as an iterator.
func (n *node) Next() *node {
	return successor(n)
}

// successor returns the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

// getColor gets color of the node.
func getColor(n *node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

// minimum finds the minimum node of subtree n.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}
