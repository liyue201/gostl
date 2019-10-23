package main

import (
	"fmt"
	"github.com/liyue201/gostl/containers/list"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n",)
}
