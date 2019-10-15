package priority_queue

import "container/heap"

type CmpFun func(a, b interface{}) bool

type ItemHolder struct {
	items  []interface{}
	cmpFun CmpFun 
}
 
func (this *ItemHolder) Push(item interface{}) {
	this.items = append(this.items, item)
}
 
func (this *ItemHolder) Pop() interface{} {
	if len(this.items) == 0 {
		return nil
	}
	item := this.items[this.Len()-1]
	this.items = this.items[:this.Len()-1]
	return item
}

func (this *ItemHolder) top() interface{} {
	if len(this.items) == 0 {
		return nil
	}
	return this.items[0]
}

func (this *ItemHolder) Size() int {
	return len(this.items)
}

func (this *ItemHolder) Len() int {
	return len(this.items)
}

func (this *ItemHolder) Less(i, j int) bool {
	return this.cmpFun(this.items[i], this.items[j])
}

func (this *ItemHolder) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

type PriorityQueue struct {
	holder *ItemHolder
}

func New(cmp CmpFun) *PriorityQueue {
	holder := &ItemHolder{
		items:  make([]interface{}, 0, 0),
		cmpFun: cmp,
	}
	return &PriorityQueue{
		holder: holder,
	}
}

func (this *PriorityQueue) Push(item interface{}) {
	heap.Push(this.holder, item)
}

func (this *PriorityQueue) Pop() interface{} {
	return heap.Pop(this.holder)
}

func (this *PriorityQueue) Top() interface{} {
	return this.holder.top()
}

func (this *PriorityQueue) Empty() bool {
	if this.holder.Size() == 0 {
		return true
	}
	return false
}
