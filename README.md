# GoSTL

GoSTL is a data structure and algorithm library for go,designed to provide functions similar to C++ STL,but more powerful.

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
    
 ## Examples

 ### <a name="slice">slice</a>
this is an example of how to create an IntSlice,and sort it.
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
```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/simple_list"
)

func main() {
	l := simple_list.New()
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
```go
package main

import (
	"fmt"
	list "github.com/liyue201/gostl/ds/list/bid_list"
)

func main() {
	l := list.New()
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

queue based on an other data structure which is a implementation of interface container.Container, such as deque or list.

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


// thread-save
func example3() {
	fmt.Printf("example3:\n")

	s := queue.New(queue.WithThreadSave())
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
priority_queue based on  package `container/heap` in go standard library
```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/priority_queue"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	q := priority_queue.New(priority_queue.WithComparator(comparator.Reverse(comparator.BuiltinTypeComparator)),
		priority_queue.WithThreadSave())
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
stack based on an other data structure which is a implementation of interface container.Container, such as deque or list.
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

// thread-save
func example3() {
	fmt.Printf("example3:\n")

	s := stack.New(stack.WithThreadSave())
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

```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/map"
)

func main() {
	m := treemap.New(treemap.WithThreadSave())

	m.Insert("a", "aaa")
	m.Insert("b", "bbb")

	fmt.Printf("a = %v\n", m.Get("a"))
	fmt.Printf("b = %v\n", m.Get("b"))
	
	m.Erase("b")
}
```

### <a name="set">set</a>
```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/set"
)

func main()  {
	s := set.New(set.WithThreadSave())
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
```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/bloomfilter"
)

func main() {
	filter := bloom.New(100, 4, bloom.WithThreadSave())
	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
```

### <a name="hamt">hamt</a>
```go
package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/hamt"
)

func main() {
	h := hamt.New()
	// h := hamt.New(hamt.WithThreadSave())
	key := []byte("aaaaa")
	val := "bbbbbbbbbbbbb"

	h.Insert(key, val)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))

	h.Erase(key)
	fmt.Printf("%v = %v\n", string(key), h.Get(key))
}
```

### <a name="ketama">ketama</a>
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