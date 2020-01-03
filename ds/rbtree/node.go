package rbtree

type Color bool

const (
	RED   = false
	BLACK = true
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  Color
	key    interface{}
	value  interface{}
}

func (n *Node) Key() interface{} {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) SetValue(val interface{}) {
	n.value = val
}

// Next returns the Node's successor as an iterator.
func (n *Node) Next() *Node {
	return successor(n)
}

// Prev returns the Node's predecessor as an iterator.
func (n *Node) Prev() *Node {
	return presuccessor(n)
}

// successor returns the successor of the Node
func successor(x *Node) *Node {
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

// presuccessor returns the presuccessor of the Node
func presuccessor(x *Node) *Node {
	if x.left != nil {
		return maximum(x.left)
	}
	if x.parent != nil {
		if x.parent.right == x {
			return x.parent
		}
		for x.parent != nil && x.parent.left == x {
			x = x.parent
		}
		return x.parent
	}
	return nil
}

// minimum finds the minimum Node of subtree n.
func minimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// maximum finds the maximum Node of subtree n.
func maximum(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}
