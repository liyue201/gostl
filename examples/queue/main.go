package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/queue"
)

func main() {
	q := queue.New()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
