package queue

import "github.com/liyue201/gostl/ds/deque"

type Queue struct {
	dq *deque.Deque
}

func New() *Queue {
	return &Queue{
		dq: deque.New(),
	}
}
func (this *Queue) Size() int {
	return this.dq.Size()
}

func (this *Queue) Empty() bool {
	return this.dq.Empty()
}

func (this *Queue) Push(value interface{}) {
	this.dq.PushBack(value)
}

func (this *Queue) Front() interface{} {
	return this.dq.Front()
}

func (this *Queue) Back() interface{} {
	return this.dq.Back()
}

func (this *Queue) Pop() interface{} {
	return this.dq.PopFront()
}

func (this *Queue) Clear() {
	this.dq.Clear()
}

func (this *Queue) String() string {
	return this.dq.String()
}
