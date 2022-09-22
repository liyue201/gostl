package slice

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceWrapper(t *testing.T) {
	type User struct {
		age  int
		name string
	}

	a := make([]*User, 0)
	a = append(a, &User{age: 2, name: "zzz"})
	a = append(a, &User{age: 5, name: "nnn"})
	a = append(a, &User{age: 2, name: "aaa"})

	sw := NewSliceWrapper(a)

	assert.Equal(t, 3, sw.Len())

	sort.Sort[*User](sw.Begin(), sw.End(), func(a, b any) int {
		ua := a.(*User)
		ub := b.(*User)
		if ua.age < ub.age {
			return -1
		}
		if ua.age > ub.age {
			return 1
		}
		if ua.name < ub.name {
			return -1
		}
		if ua.name > ub.name {
			return 1
		}
		return 0
	})
	for i := range a {
		t.Logf("%+v\n", a[i])
		assert.Equal(t, a[i], sw.At(i))
		if i > 0 {
			assert.LessOrEqual(t, a[i-1].age, a[i].age)
			if a[i-1].age == a[i].age {
				assert.LessOrEqual(t, a[i-1].name, a[i].name)
			}
		}
	}
}

func TestIntSlice(t *testing.T) {
	a := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	sliceA := NewSliceWrapper(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, 100)
	assert.EqualValues(t, 100, sliceA.At(8))
}
