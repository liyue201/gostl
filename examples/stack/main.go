package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/stack"
	"sync"
	"time"
)

func example1() {
	fmt.Printf("example1:\n")

	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.Empty() {
		fmt.Printf("%v\n", s.Pop())
	}
}

// based on list
func example2() {
	fmt.Printf("example2:\n")

	s := stack.New(stack.WithListContainer())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.Empty() {
		fmt.Printf("%v\n", s.Pop())
	}
}

// thread-save
func example3() {
	fmt.Printf("example3:\n")

	s := stack.New(stack.WithThreadSafe())
	sw := sync.WaitGroup{}
	sw.Add(2)
	go func() {
		defer sw.Done()
		for i := 0; i < 10; i++ {
			s.Push(i)
			time.Sleep(time.Microsecond * 100)
		}
	}()

	go func() {
		defer sw.Done()
		for i := 0; i < 10; {
			if !s.Empty() {
				val := s.Pop()
				fmt.Printf("%v\n", val)
				i++
			} else {
				time.Sleep(time.Microsecond * 100)
			}
		}
	}()
	sw.Wait()
}

func main() {
	example1()
	example2()
	example3()
}
