package slice

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/stretchr/testify/assert"
	"reflect"
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

	sw := NewSliceWrapper(a, reflect.TypeOf(&User{}))

	assert.Equal(t, 3, sw.Len())

	sort.Sort(sw.Begin(), sw.End(), func(a, b interface{}) int {
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
