package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/bidlist"
	"github.com/liyue201/gostl/ds/list/simplelist"
)

func simpleList() {
	l := simplelist.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.FrontNode(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n===============\n")
}

func bidList() {
	l := bidlist.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.FrontNode(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n")

	for n := l.BackNode(); n != nil; n = n.Prev() {
		fmt.Printf("%v ", n.Value)
	}
}

func main() {
	simpleList()
	bidList()
}
