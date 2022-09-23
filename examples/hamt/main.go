package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/hamt"
)

func main() {
	h := hamt.New[string](hamt.WithGoroutineSafe())
	key := []byte("aaaaa")
	val := "bbbbbbbbbbbbb"

	h.Insert(key, val)
	v, _ := h.Get(key)
	fmt.Printf("%v = %v\n", string(key), v)

	h.Erase(key)
	v, _ = h.Get(key)
	fmt.Printf("%v = %v\n", string(key), v)
}
