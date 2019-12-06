package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/skiplist"
)

func main() {
	list := skiplist.New(skiplist.WithMaxLevel(15))
	list.Insert("aaa", "1111")
	list.Insert("bbb", "2222")
	fmt.Printf("aaa = %v\n", list.Get("aaa"))
	fmt.Printf("aaa = %v\n\n", list.Get("bbb"))

	list.Traversal(func(key, value interface{}) bool {
		fmt.Printf("key:%v value:%v\n", key, value)
		return true
	})

	list.Remove("aaa")
}
