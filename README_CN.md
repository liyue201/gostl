# GoSTL

[![GoDoc](https://godoc.org/github.com/liyue201/gostl?status.svg)](https://godoc.org/github.com/liyue201/gostl)
[![Goreportcard](https://goreportcard.com/badge/github.com/liyue201/gostl)](https://goreportcard.com/report/github.com/liyue201/gostl)
[![Build Status](https://travis-ci.org/liyue201/gostl.svg?branch=master)](https://travis-ci.org/liyue201/gostl)
[![Coverall](https://coveralls.io/repos/github/liyue201/gostl/badge.svg?branch=master)](https://coveralls.io/github/liyue201/gostl)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](/LICENSE)
[![View examples](https://img.shields.io/badge/learn-examples-brightgreen.svg)](/examples)

[English](./README.md) | 简体中文

GoSTL是一个go语言数据结构和算法库，类似C++的STL，但功能更强大。结合go语言的特点，大部分数据结构都实现了协程安全，可以在创建对象的时候通过配置参数指定是否开启。

## 功能列表
- 数据结构
    - [切片（slice）](#slice)
    - [数组（array）](#array)
    - [向量（vector）](#vector)
    - [链表（list）](#list)
    - [双端队列（deque）](#deque)
    - [队列（queue）](#queue)
    - [优先队列（priority_queue）](#priority_queue)
    - [栈（stack）](#stack)
    - [红黑树（rbtree）](#rbtree)
    - [映射（map/multimap）](#map)
    - [集合（set/multiset）](#set)
    - [位映射（bitmap）](#bitmap)
    - [布隆过滤器（bloom_filter）](#bloom_filter)
    - [哈希数组映射字典树（hash_array_mapped_trie）](#hamt)
    - [一致性哈希（ketama）](#ketama)
    - [跳表（skiplist）](#skliplist)
- 算法
    - [快排（sort）](#sort)
    - [稳定排序（stable_sort）](#sort)
    - [二分查找（binary_search）](#sort)
    - [二分查找第一个元素的位置（lower_bound）](#sort)
    - [二分查找第一个大于该元素的位置（upper_bound）](#sort)
    - [下一个排列组合（next_permutation）](#next_permutation)
    - [第n个元素（nth_element)](#nth_element)
    - [交换（swap）](#algo_op)
    - [翻转（reverse）](#algo_op)
    - [统计（count/count_if）](#algo_op_const)
    - [查找（find/find_if）](#algo_op_const)
    - [最大最小值（min_element/max_element）](#algo_op_const)
      
## 例子

### <a name="slice">切片（slice）</a>
这个库中的切片是对go原生切片的定义包装。  


 ```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/slice"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  a := make([]int, 0)
  a = append(a, 2)
  a = append(a, 1)
  a = append(a, 3)
  fmt.Printf("%v\n", a)

  wa := slice.NewSliceWrapper(a)

  // sort in ascending
  sort.Sort[int](wa.Begin(), wa.End(), comparator.IntComparator)
  fmt.Printf("%v\n", a)

  // sort in descending
  sort.Sort[int](wa.Begin(), wa.End(), comparator.Reverse(comparator.IntComparator))
  fmt.Printf("%v\n", a)
}

 ```
 
### <a name="array">数组（array）</a>
数组是一种一旦创建长度就固定的数据结构，支持随机访问和迭代器访问。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/array"
)

func main() {
  a := array.New[int](5)
  for i := 0; i < a.Size(); i++ {
    a.Set(i, i+1)
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
向量是一种大小可以自动伸缩的数据结构，内部使用切片实现。支持随机访问和迭代器访问。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/vector"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  v := vector.New[int]()
  v.PushBack(1)
  v.PushBack(2)
  v.PushBack(3)
  for i := 0; i < v.Size(); i++ {
    fmt.Printf("%v ", v.At(i))
  }
  fmt.Printf("\n")

  // sort in descending
  sort.Sort[int](v.Begin(), v.End(), comparator.Reverse(comparator.IntComparator))
  for iter := v.Begin(); iter.IsValid(); iter.Next() {
    fmt.Printf("%v ", iter.Value())
  }
}

```


### <a name="list">链表（list）</a>

- 简单链表  
简单链表是一种单向链表，支持从头部和尾部插入数据，只支持从头部遍历数据。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/list/simplelist"
)

func main() {
  l := simplelist.New[int]()
  l.PushBack(1)
  l.PushFront(2)
  l.PushFront(3)
  l.PushBack(4)
  for n := l.FrontNode(); n != nil; n = n.Next() {
    fmt.Printf("%v ", n.Value)
  }
  fmt.Printf("\n===============\n")
}
```
  
- 双向链表  
双向链表，支持从头部和尾部插入数据，支持从头部和尾部遍历数据。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/list/bidlist"
)
func main() {
  l := bidlist.New[int]()
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
```

### <a name="deque">双端队列（deque）</a>
双端队列支持从头部和尾部高效的插入数据，支持随机访问和迭代器访问。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/deque"
  "github.com/liyue201/gostl/utils/comparator"
  "math/rand"
)

func main() {
  q := deque.New[int]()
  for i := 0; i < 100; i++ {
    r := rand.Int() % 100
    q.PushBack(r)
    q.PushFront(r)
  }
  fmt.Printf("%v\n", q)

  sort.Sort[int](q.Begin(), q.End(), comparator.IntComparator)
  fmt.Printf("%v\n", q)

  for !q.Empty() {
    r := rand.Int() % q.Size()
    q.EraseAt(r)
  }
  fmt.Printf("%v\n", q)
}

```

### <a name="queue">队列（queue）</a>
队列是一种先进先出的数据结构，底层使用双端队列或者链表作为容器，默认使用双端队列，若想使用链表，可以在创建对象时使用queue.WithListContainer()参数。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/queue"
)

func main() {
  q := queue.New[int]()
  for i := 0; i < 5; i++ {
    q.Push(i)
  }
  for !q.Empty() {
    fmt.Printf("%v\n", q.Pop())
  }
}
```

### <a name="priority_queue">优先队列（priority_queue）</a>
优先级队列是一种类似于队列的抽象数据类型，每个元素都有一些与之关联的优先级值。 优先级队列中元素的优先级决定了移除元素的顺序。 

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/priorityqueue"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  q := priorityqueue.New[int](comparator.Reverse(comparator.IntComparator),
    priorityqueue.WithGoroutineSafe())
  q.Push(4)
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

### <a name="stack">栈（stack）</a>
栈是一种后进先出的数据结构，底层使用双端队列或者链表作为容器，默认使用双端队列，若想使用链表，可以在创建对象时使用queue.WithListContainer()参数。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/stack"
)

func main() {
  s := stack.New[int]()
  s.Push(1)
  s.Push(2)
  s.Push(3)
  for !s.Empty() {
    fmt.Printf("%v\n", s.Pop())
  }
}
```

### <a name="rbtree">红黑树（rbtree）</a>
红黑树是一种平衡二叉排序树，用于高效的插入和查找数据。

```go
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


```

### <a name="map">映射（map）</a>
映射底层使用红黑树实现，支持按key顺序迭代访问，有别于go原生的map类型（go原生的map底层是哈希，不支持按key顺序迭代访问）。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/map"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  m := treemap.New[string, string](comparator.StringComparator, treemap.WithGoroutineSafe())

  m.Insert("a", "aaa")
  m.Insert("b", "bbb")

  a, _ := m.Get("a")
  b, _ := m.Get("b")
  fmt.Printf("a = %v\n", a)
  fmt.Printf("b = %v\n", b)

  m.Erase("b")
}

```

### <a name="set">集合（set）</a>
集合底层使用红黑树实现，支持线程安全。支持集合的基本运算，如求并集，交集，差集。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/set"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  s := set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
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

### <a name="bitmap">位映射（bitmap）</a>
位映射用于快速标记和查找一个非负整数是否集合中。相对于map或数组占用内存空间更小。

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

### <a name="bloom_filter">布隆过滤器（bloom_filter）</a>
布隆过滤器用来快速判断数据是否在集合中，底层使用bitmap实现，相对于map占用内存空间更小。缺点是不支持删除和有一定的错误率。支持线程安全。支持数据导出和通过导出的数据重新构建。

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

### <a name="hamt">哈希数组映射字典树（hash_array_mapped_trie）</a>
哈希数组映射字典树相对于传统哈希（开放地址法或链表法哈希），出现哈希冲突的概率更小，空间利用率更高。扩容缩容性时间复杂度低。支持线程安全。

```go
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

```

### <a name="ketama">一致性哈希（ketama）</a>
一致性哈希ketama算法，使用64位的哈希函数和map存储，出现冲突的概率更小。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/ketama"
)

func main() {
  k := ketama.New()
  k.Add("1.2.3.3")
  k.Add("2.4.5.6")
  k.Add("5.5.5.1")

  for i := 0; i < 10; i++ {
    node, _ := k.Get(fmt.Sprintf("%d", i))
    fmt.Printf("%v\n", node)
  }
  k.Remove("2.4.5.6")
}


```
### <a name="skliplist">跳表（skliplist）</a>
跳表是一种通过以空间换时间来实现快速查找的数据结构。支持线程安全。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/ds/skiplist"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  list := skiplist.New[string, string](comparator.StringComparator, skiplist.WithMaxLevel(15))
  list.Insert("aaa", "1111")
  list.Insert("bbb", "2222")
  v1, _ := list.Get("aaa")
  v2, _ := list.Get("bbb")
  fmt.Printf("aaa = %v\n", v1)
  fmt.Printf("bbb = %v\n", v2)

  list.Traversal(func(key, value string) bool {
    fmt.Printf("key:%v value:%v\n", key, value)
    return true
  })

  list.Remove("aaa")
}

```

### <a name="sort">排序、稳定排序、二分查找</a>
- Sort: 内部使用的是快速排序算法。 
- Stable: 稳定排序，内部使用归并排序。    
- BinarySearch: 通过二分查找，判断一个元素是否在迭代器范围中。  
- LowerBound: 通过二分查找，找到第一个等于该元素的数据返回该迭代器。  
- UpperBound：  通过二分查找，找到第一个大于该元素的数据返回该迭代器。  

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/slice"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  a := make([]string, 0)
  a = append(a, "bbbb")
  a = append(a, "ccc")
  a = append(a, "aaaa")
  a = append(a, "bbbb")
  a = append(a, "bb")

  sliceA := slice.NewSliceWrapper(a)

  ////Sort in ascending order
  sort.Sort[string](sliceA.Begin(), sliceA.End(), comparator.OrderedTypeCmp[string])

  sort.Stable[string](sliceA.Begin(), sliceA.End(), comparator.StringComparator)
  fmt.Printf("%v\n", a)

  if sort.BinarySearch[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator) {
    fmt.Printf("BinarySearch: found bbbb\n")
  }

  iter := sort.LowerBound[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator)
  if iter.IsValid() {
    fmt.Printf("LowerBound bbbb: %v\n", iter.Value())
  }
  iter = sort.UpperBound[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator)
  if iter.IsValid() {
    fmt.Printf("UpperBound bbbb: %v\n", iter.Value())
  }
  //Sort in descending order
  sort.Sort[string](sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.StringComparator))
  //sort.Stable[string](sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.StringComparator))
  fmt.Printf("%v\n", a)
}

```

### <a name="next_permutation">下个排列组合（next_permutation）</a>
这个函数修改迭代器范围内的数据为下一个排列组合。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/slice"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  a := make([]int, 0)
  for i := 1; i <= 3; i++ {
    a = append(a, i)
  }
  wa := slice.NewSliceWrapper(a)
  fmt.Println("NextPermutation")
  for {
    fmt.Printf("%v\n", a)
    if !sort.NextPermutation[int](wa.Begin(), wa.End(), comparator.IntComparator) {
      break
    }
  }
  fmt.Println("PrePermutation")
  for {
    fmt.Printf("%v\n", a)
    if !sort.NextPermutation[int](wa.Begin(), wa.End(), comparator.Reverse(comparator.IntComparator)) {
      break
    }
  }
}

```

### <a name="nth_element">第n个元素（nth_element）</a>
将迭代器范围内的第n个元素放在n的位置，并将小于或等于它的元素放在左边，大于或等于它的元素放在右边。

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm/sort"
  "github.com/liyue201/gostl/ds/deque"
  "github.com/liyue201/gostl/utils/comparator"
)

func main() {
  a := deque.New[int]()
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
  sort.NthElement[int](a.Begin(), a.End(), 3, comparator.IntComparator)
  fmt.Printf("%v\n", a.At(3))
  fmt.Printf("%v\n", a)
}
```

### <a name="algo_op"> 交换/翻转 </a>
- 交换（swap):  交换两个迭代器的值  
- 翻转(reverse): 翻转两个迭代器区间范围内的值  

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm"
  "github.com/liyue201/gostl/ds/deque"
)

func main() {
  a := deque.New[int]()
  for i := 0; i < 9; i++ {
    a.PushBack(i)
  }
  fmt.Printf("%v\n", a)

  algorithm.Swap[int](a.First(), a.Last())
  fmt.Printf("%v\n", a)

  algorithm.Reverse[int](a.Begin(), a.End())
  fmt.Printf("%v\n", a)
}

```

### <a name="algo_op_const"> count/count_if/find/find_if/min_element/max_element </a>
- Count : 在迭代器区间内统计等于指定值的数量
- CountIf： 在迭代器区间内统计等于满足函数f的数量
- Find：在迭代器区间内找到第一个等于指定值的元素，返回其迭代器
- FindIf：在迭代器区间内找到第一个满足函数f的元素，返回其迭代器

```go
package main

import (
  "fmt"
  "github.com/liyue201/gostl/algorithm"
  "github.com/liyue201/gostl/ds/deque"
  "github.com/liyue201/gostl/utils/comparator"
  "github.com/liyue201/gostl/utils/iterator"
)

func isEven(iter iterator.ConstIterator[int]) bool {
  return iter.Value()%2 == 0
}

func greaterThan5(iter iterator.ConstIterator[int]) bool {
  return iter.Value() > 5
}

func main() {
  a := deque.New[int]()
  for i := 0; i < 10; i++ {
    a.PushBack(i)
  }
  for i := 0; i < 5; i++ {
    a.PushBack(i)
  }
  fmt.Printf("%v\n", a)

  fmt.Printf("Count 2: %v\n", algorithm.Count[int](a.Begin(), a.End(), 2, comparator.IntComparator))

  fmt.Printf("Count 2: %v\n", algorithm.CountIf[int](a.Begin(), a.End(), isEven))

  iter := algorithm.Find[int](a.Begin(), a.End(), 2, comparator.IntComparator)
  if !iter.Equal(a.End()) {
    fmt.Printf("Fund %v\n", iter.Value())
  }
  iter = algorithm.FindIf[int](a.Begin(), a.End(), greaterThan5)
  if !iter.Equal(a.End()) {
    fmt.Printf("FindIf greaterThan5 : %v\n", iter.Value())
  }
  iter = algorithm.MaxElement[int](a.Begin(), a.End(), comparator.IntComparator)
  if !iter.Equal(a.End()) {
    fmt.Printf("largest value : %v\n", iter.Value())
  }
  iter = algorithm.MinElement[int](a.Begin(), a.End(), comparator.IntComparator)
  if !iter.Equal(a.End()) {
    fmt.Printf("largest value : %v\n", iter.Value())
  }
}

```
