package main

import (
	"fmt"
	"github.com/liyue201/gostl/containers/queue"
)

func main() {
	q := queue.New(10)
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
