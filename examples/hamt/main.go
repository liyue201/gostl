package main

import (
	"fmt"
	"github.com/liyue201/gostl/containers/hamt"
)

func main() {
	h := hamt.New()
	key := []byte("aaaaa")
	val := "bbbbbbbbbbbbb"

	h.Insert(key, val)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))

	h.Insert(key, "hhhh")
	fmt.Printf("%v = %v\n", string(key), h.Get(key))

	h.Erase(key)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))
}
