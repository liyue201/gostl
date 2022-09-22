package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
)

func main() {
	tree := rbtree.New()
	tree.Insert(1, "aaa")
	tree.Insert(5, "bbb")
	tree.Insert(3, "ccc")
	fmt.Printf("find %v returns %v\n", 5, tree.Find(5))

	tree.Traversal(func(key, value any) bool {
		fmt.Printf("%v : %v\n", key, value)
		return true
	})
	tree.Delete(tree.FindNode(3))
}
