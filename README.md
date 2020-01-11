# GoSTL

[![GoDoc](https://godoc.org/github.com/liyue201/gostl?status.svg)](https://godoc.org/github.com/liyue201/gostl)
[![](https://goreportcard.com/badge/github.com/liyue201/gostl)](https://goreportcard.com/report/github.com/liyue201/gostl)
[![Build Status](https://travis-ci.org/liyue201/gostl.svg?branch=master)](https://travis-ci.org/liyue201/gostl)
[![](https://coveralls.io/repos/github/liyue201/gostl/badge.svg?branch=master)](https://coveralls.io/github/liyue201/gostl)
[![licence](https://img.shields.io/badge/license-GPLv3-brightgreen.svg)]()
[![view examples](https://img.shields.io/badge/learn-examples-brightgreen.svg)](/examples)

English | [简体中文](./README_CN.md)

## Introduction
GoSTL is a data structure and algorithm library for go,designed to provide functions similar to C++ STL,but more powerful.Combined with the characteristics of go language, most of the data structures have realized goroutine-safe. When creating objects, you can specify whether to turn it on or not through configuration parameters.


## Function list
- data structure
    - [slice](#slice)
    - [array](#array)
    - [vector](#vector)
    - [list](#list)
    - [deque](#deque)
    - [queue](#queue)
    - [priority_queue](#priority_queue)
    - [stack](#stack)
    - [rbtree(red_black_tree)](#rbtree)
    - [map/multimap](#map)
    - [set/multiset](#set)
    - [bitmap](#bitmap)
    - [bloom_filter](#bloom_filter)
    - [hamt(hash_array_map_tree)](#hamt)
    - [ketama](#ketama)
    - [skiplist](#skliplist)
- algorithm
    - [sort(quick_sort)](#sort)
    - [stable_sort(merge_sort)](#sort)
    - [binary_search](#sort)
    - [lower_bound](#sort)
    - [upper_bound](#sort)
    - [next_permutation](#next_permutation)
    - [nth_element](#nth_element)
    - [swap](#algo_op)
    - [reverse](#algo_op)
    - [count/count_if](#algo_op_const)
    - [find/find_if](#algo_op_const)
    
 ## Examples

 ### <a name="slice">slice</a>
The slice in this library is a redefinition of go native slice.
Here is an example of how to turn a go native slice into an Intslice and sort it. 

 ```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	a := slice.IntSlice(make([]int, 0))
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 3)
	fmt.Printf("%v\n", a)
	
	// sort in ascending
	sort.Sort(a.Begin(), a.End())
	fmt.Printf("%v\n", a)
	
	// sort in descending
	sort.Sort(a.Begin(), a.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v\n", a)
}
 ```
 
### <a name="array">array</a>
Array is a data structure with fixed length once it is created, which supports random access and iterator access.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/array"
)

func main() {
	a := array.New(5)
	for i := 0; i < a.Size(); i++ {
		a.Set(i, i + 1)
	}
	for i := 0; i < a.Size(); i++ {
		fmt.Printf("%v ", a.At(i))
	}

	fmt.Printf("\n")
	for iter := a.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
}

```

### <a name="vector">vector</a>
Vector is a kind of data structure whose size can be automatically expanded, which is realized by slice internally. Supports random access and iterator access.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/vector"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	v := vector.New()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	for i := 0; i < v.Size(); i++ {
		fmt.Printf("%v ", v.At(i))
	}
	fmt.Printf("\n")

	// sort in descending
	sort.Sort(v.Begin(), v.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
}
```


### <a name="list">list</a>
- simple list  
Simple list is a one directional list, which supports inserting data from the head and tail, and only traversing data from the head.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/simplelist"
)

func main() {
	l := simplelist.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.FrontNode(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
}

```
  
- bidirectional list  
Bidirectional list supports inserting data from the head and tail, and traversing data from the head and tail.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/bidlist"
)

func main() {
	l := bidlist.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.FrontNode(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
	fmt.Printf("\n", )

	for n := l.BackNode(); n != nil; n = n.Prev() {
		fmt.Printf("%v ", n.Value)
	}
}
```

### <a name="deque">deque</a>
Deque supports efficient data insertion from the head and tail, random access and iterator access.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	q := deque.New()
	q.PushBack(2)
	q.PushFront(1)
	q.PushBack(3)
	q.PushFront(4)
	fmt.Printf("%v\n", q)

	sort.Sort(q.Begin(), q.End())
	fmt.Printf("%v\n", q)
	fmt.Printf("%v\n", q.PopBack())
	fmt.Printf("%v\n", q.PopFront())
	fmt.Printf("%v\n", q)
}

```

### <a name="queue">queue</a>
Queue is a first-in-first-out data structure. The bottom layer uses the deque or list as the container. By default, the deque is used. If you want to use the list, you can use the `queue.WithListContainer()` parameter when creating an object. Goroutine safety is supported.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/queue"
	"sync"
	"time"
)

func example1()  {
	fmt.Printf("example1:\n")
	q := queue.New()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}

//  based on list
func example2()  {
	fmt.Printf("example2:\n")
	q := queue.New(queue.WithListContainer())
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}


// goroutine-save
func example3() {
	fmt.Printf("example3:\n")

	s := queue.New(queue.WithGoroutineSafe())
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
```

### <a name="priority_queue">priority_queue</a>
priority_queue is based on the  `container/heap`  package of go standard library, which supports goroutine safety.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	q := priorityqueue.New(priorityqueue.WithComparator(comparator.Reverse(comparator.BuiltinTypeComparator)),
		priorityqueue.WithGoroutineSafe())
	q.Push(5)
	q.Push(13)
	q.Push(7)
	q.Push(9)
	q.Push(0)
	q.Push(88)
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
```

### <a name="stack">stack</a>
Stack is a kind of last-in-first-out data structure. The bottom layer uses the deque or list as the container. By default, the deque is used. If you want to use the list, you can use the `queue.WithListContainer()` parameter when creating an object.  Goroutine safety is supported.

```go
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

// goroutine-save
func example3() {
	fmt.Printf("example3:\n")

	s := stack.New(stack.WithGoroutineSafe())
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
```

### <a name="rbtree">rbtree</a>
Red black tree is a balanced binary sort tree, which is used to insert and find data efficiently.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
)

func main()  {
	tree := rbtree.New()
	tree.Insert(1, "aaa")
	tree.Insert(5, "bbb")
	tree.Insert(3, "ccc")
	fmt.Printf("find %v returns %v\n",5, tree.Find(5))

	tree.Traversal(func(key, value interface{}) bool {
		fmt.Printf("%v : %v\n", key, value)
		return true
	})
	tree.Delete(tree.FindNode(3))
}

```

### <a name="map">map</a>
The Map bottom layer is implemented by using red black tree, and supports iterative access in key order, which is different from the go native map type (the go native map bottom layer is hash, and does not support iterative access in key order). Goroutine safety is supported.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/map"
)

func main() {
	m := treemap.New(treemap.WithGoroutineSafe())

	m.Insert("a", "aaa")
	m.Insert("b", "bbb")

	fmt.Printf("a = %v\n", m.Get("a"))
	fmt.Printf("b = %v\n", m.Get("b"))
	
	m.Erase("b")
}
```

### <a name="set">set</a>
The Set bottom layer is implemented by red black tree, which supports goroutine safety. Support basic operations of set, such as union, intersection and difference. Goroutine safety is supported.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/set"
)

func main()  {
	s := set.New(set.WithGoroutineSafe())
	s.Insert(1)
	s.Insert(5)
	s.Insert(3)
	s.Insert(4)
	s.Insert(2)

	s.Erase(4)

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v\n", iter.Value())
	}

	fmt.Printf("%v\n", s.Contains(3))
	fmt.Printf("%v\n", s.Contains(10))
}
```

### <a name="bitmap">bitmap</a>
Bitmap is used to quickly mark and find whether a non negative integer is in a set. It takes up less memory than map or array.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/bitmap"
)

func main() {
	bm := bitmap.New(1000)
	bm.Set(6)
	bm.Set(10)

	fmt.Printf("%v\n", bm.IsSet(5))
	fmt.Printf("%v\n", bm.IsSet(6))
	bm.Unset(6)
	fmt.Printf("%v\n", bm.IsSet(6))
}
```
### <a name="bloom_filter">bloom_filter</a>
Boomfilter is used to quickly determine whether the data is in the collection. The bottom layer is implemented with bitmap, which uses less memory than map. The disadvantage is that it does not support deletion and has a certain error rate. Goroutine safety is supported , supports data export and reconstruction through exported data.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/bloomfilter"
)

func main() {
	filter := bloom.New(100, 4, bloom.WithGoroutineSafe())
	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
```

### <a name="hamt">hamt</a>
Compared with the traditional hash (open address method or linked list method hash), hamt has lower probability of hash conflict and higher space utilization. The time complexity of capacity expansion is low. Goroutine safety is supported.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/hamt"
)

func main() {
	h := hamt.New()
	// h := hamt.New(hamt.WithGoroutineSafe())
	key := []byte("aaaaa")
	val := "bbbbbbbbbbbbb"

	h.Insert(key, val)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))

	h.Erase(key)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))
}
```

### <a name="ketama">ketama</a>
Consistent hash Ketama algorithm, using 64 bit hash function and map storage, has less conflict probability. Goroutine safety is supported.

```go
package main

import (
	"github.com/liyue201/gostl/ds/ketama"
	"fmt"
)

func main() {
	k := ketama.New()
	k.Add("1.2.3.3")
	k.Add("2.4.5.6")
	k.Add("5.5.5.1")
	node, _ := k.Get("aaa")
    fmt.Printf("%v\n", node)
	k.Remove("2.4.5.6")
}

```
### <a name="skliplist">skliplist</a>
Skliplist is a kind of data structure which can search quickly by exchanging space for time. Goroutine safety is supported.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/skiplist"
)

func main()  {
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
```

### <a name="sort">sort</a>
Sort: quick sort algorithm is used internally.  
Stable: stable sorting. Merge sorting is used internally.  
Binarysearch: determine whether an element is in the scope of iterator by binary search.  
Lowerbound: find the first data equal to the element and return the iterator by binary search.  
Upperbound: find the first data larger than the element and return the iterator by binary search.  

```go
package main

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/ds/slice"
	"fmt"
)

func main()  {
	a := make([]string, 0)
	a = append(a, "bbbb")
	a = append(a, "ccc")
	a = append(a, "aaaa")
	a = append(a, "bbbb")
	a = append(a, "bb")

	sliceA := slice.StringSlice(a)

	////Sort in ascending order
	sort.Sort(sliceA.Begin(), sliceA.End())
	//sort.Stable(sliceA.Begin(), sliceA.End())
	fmt.Printf("%v\n", a)

	if sort.BinarySearch(sliceA.Begin(), sliceA.End(), "bbbb") {
		fmt.Printf("BinarySearch: found bbbb\n")
	}

	iter := sort.LowerBound(sliceA.Begin(), sliceA.End(), "bbbb")
	if iter.IsValid() {
		fmt.Printf("LowerBound bbbb: %v\n", iter.Value())
	}
	iter = sort.UpperBound(sliceA.Begin(), sliceA.End(), "bbbb")
	if iter.IsValid() {
		fmt.Printf("UpperBound bbbb: %v\n", iter.Value())
	}
	//Sort in descending order
	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	//sort.Stable(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v\n", a)
}
```

### <a name="next_permutation">next_permutation</a>
This function modifies the data in the iterator range to the next sort combination.

```go
package main

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
	"fmt"
)

func main()  {
	a := slice.IntSlice(make([]int, 0))

	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	fmt.Println("NextPermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation(a.Begin(), a.End()) {
			break
		}
	}
	fmt.Println("PrePermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation(a.Begin(), a.End(), comparator.Reverse(comparator.BuiltinTypeComparator)) {
			break
		}
	}
}
```

### <a name="nth_element">nth_element</a>
Place the nth element in the scope of the iterator in the position of N, and put the element less than or equal to it on the left, and the element greater than or equal to it on the right.

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	a := deque.New()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	a.PushBack(6)
	a.PushBack(5)
	a.PushBack(4)
	a.PushBack(3)
	a.PushBack(2)
	a.PushBack(1)
	fmt.Printf("%v\n", a)
	sort.NthElement(a.Begin(), a.End(), 3)
	fmt.Printf("%v\n", a.At(3))
	fmt.Printf("%v\n", a)
}
```

### <a name="algo_op"> swap/reverse </a>
- swap: swap the values of two iterators
- reverse: Reverse values in the range of two iterators

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	a := deque.New()
	for i := 0; i < 9; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	algorithm.Swap(a.First(), a.Last())
	fmt.Printf("%v\n", a)

	algorithm.Reverse(a.Begin(), a.End())
	fmt.Printf("%v\n", a)
}

```

### <a name="algo_op_const"> count/count_if/find/find_if </a>
- Count : Count the number of elements equal to the specified value in the iterator interval
- CountIf: Count the number of elements that satisfy the function f in the iterator interval
- Find: Find the first element equal to the specified value in the iterator interval and returns its iterator
- FindIf：Find the first element satisfying function f in the iterator interval and return its iterator

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/utils/iterator"
)

func isEven(iter iterator.ConstIterator) bool {
	return iter.Value().(int)%2 == 0
}

func greaterThan5(iter iterator.ConstIterator) bool {
	return iter.Value().(int) > 5
}

func main() {
	a := deque.New()
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	for i := 0; i < 5; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	fmt.Printf("Count 2: %v\n", algorithm.Count(a.Begin(), a.End(), 2))
	fmt.Printf("Count 2: %v\n", algorithm.CountIf(a.Begin(), a.End(), isEven))

	iter := algorithm.Find(a.Begin(), a.End(), 2)
	if !iter.Equal(a.End()) {
		fmt.Printf("Fund %v\n", iter.Value())
	}
	iter = algorithm.FindIf(a.Begin(), a.End(), greaterThan5)
	if !iter.Equal(a.End()) {
		fmt.Printf("FindIf greaterThan5 : %v\n", iter.Value())
	}
}

```
