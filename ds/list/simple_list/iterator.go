package simple_list

import . "github.com/liyue201/gostl/utils/iterator"

//ListIterator is an implementation Iterator
var _ Iterator = (*ListIterator)(nil)

type ListIterator struct {
	node *Node
}

func NewIterator(node *Node) *ListIterator {
	return &ListIterator{node: node}
}

func (this *ListIterator) IsValid() bool {
	return this.node != nil
}

func (this *ListIterator) Next() ConstIterator {
	if this.node != nil {
		this.node = this.node.Next()
	}
	return this
}

func (this *ListIterator) Value() interface{} {
	if this.node == nil {
		return nil
	}
	return this.node.Value
}

func (this *ListIterator) SetValue(value interface{}) error {
	if this.node != nil {
		this.node.Value = value
	}
	return nil
}

func (this *ListIterator) Clone() ConstIterator {
	return NewIterator(this.node)
}

func (this *ListIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*ListIterator)
	if !ok {
		return false
	}
	if otherIter.node == this.node {
		return true
	}
	return false
}
