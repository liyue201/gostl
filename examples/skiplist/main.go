package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/skiplist"
)

func main() {
	list := skiplist.New[string, string](skiplist.WithMaxLevel(15))
	list.Insert("aaa", "1111")
	list.Insert("bbb", "2222")
	v1, _ := list.Get("aaa")
	v2, _ := list.Get("bbb")
	fmt.Printf("aaa = %v\n", v1)
	fmt.Printf("aaa = %v\n\n", v2)

	list.Traversal(func(key, value string) bool {
		fmt.Printf("key:%v value:%v\n", key, value)
		return true
	})

	list.Remove("aaa")
}
