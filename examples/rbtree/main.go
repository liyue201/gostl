package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	tree := rbtree.New[int, string](comparator.IntComparator)
	tree.Insert(1, "aaa")
	tree.Insert(5, "bbb")
	tree.Insert(3, "ccc")
	v, _ := tree.Find(5)
	fmt.Printf("find %v returns %v\n", 5, v)

	tree.Traversal(func(key int, value string) bool {
		fmt.Printf("%v : %v\n", key, value)
		return true
	})
	tree.Delete(tree.FindNode(3))
}
