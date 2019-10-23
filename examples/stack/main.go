package main

import (
	"github.com/liyue201/gostl/containers/stack"
	"fmt"
)

func main()  {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.Empty()  {
		fmt.Printf("%v\n", s.Pop())
	}
}