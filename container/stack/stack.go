package stack

import "github.com/liyue201/gostl/container/deque"

type Stack struct {
	dq *deque.Deque
}

func New() *Stack {
	return &Stack{
		dq: deque.New(0),
	}
}

func (this *Stack) Size() int {
	return this.dq.Size()
}

func (this *Stack) Empty() bool {
	return this.dq.Empty()
}

func (this *Stack) Push(value interface{}) {
	this.dq.PushBack(value)
}

func (this *Stack) Top() interface{} {
	return this.dq.Back()
}

func (this *Stack) Pop() interface{} {
	return this.dq.PopBack()
}

func (this *Stack) Clear() {
	this.dq.Clear()
}

func (this *Stack) String() string {
	return this.dq.String()
}
