package comparator

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

// Comparator Should return a number:
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
type Comparator[T any] func(a, b T) int

func OrderedTypeCmp[T Ordered](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Reverse returns a comparator reverse to cmp
func Reverse[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return -cmp(a, b)
	}
}

// IntComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func IntComparator(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// UintComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func UintComparator(a, b uint) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int8Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int8Comparator(a, b int8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint8Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint8Comparator(a, b uint8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int16Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int16Comparator(a, b int16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint16Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint16Comparator(a, b uint16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int32Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int32Comparator(a, b int32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint32Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint32Comparator(a, b uint32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int64Comparator(a, b int64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint64Comparator(a, b uint64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Float32Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Float32Comparator(a, b float32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Float64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Float64Comparator(a, b float64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// StringComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func StringComparator(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// UintptrComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func UintptrComparator(a, b uintptr) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// BoolComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func BoolComparator(a, b bool) int {
	if a == b {
		return 0
	}
	if !a && b {
		return -1
	}
	return 1
}

// Complex64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Complex64Comparator(a, b complex64) int {
	if a == b {
		return 0
	}
	if real(a) < real(a) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}

// Complex128Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Complex128Comparator(a, b complex128) int {
	if a == b {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}
