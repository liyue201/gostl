package comparator

// Comparator Should return a number:
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
type Comparator func(a, b interface{}) int

// BuiltinTypeComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
// make sure a and b are both builtin type
func BuiltinTypeComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	switch a.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, uintptr:
		return cmpInt(a, b)
	case float32:
		if a.(float32) < b.(float32) {
			return -1
		}
	case float64:
		if a.(float64) < b.(float64) {
			return -1
		}
	case bool:
		if a.(bool) == false && b.(bool) == true {
			return -1
		}
	case string:
		if a.(string) < b.(string) {
			return -1
		}
	case complex64:
		return cmpComplex64(a.(complex64), b.(complex64))
	case complex128:
		return cmpComplex128(a.(complex128), b.(complex128))
	}
	return 1
}

func cmpInt(a, b interface{}) int {
	switch a.(type) {
	case int:
		return cmpInt64(int64(a.(int)), int64(b.(int)))
	case uint:
		return cmpUint64(uint64(a.(uint)), uint64(b.(uint)))
	case int8:
		return cmpInt64(int64(a.(int8)), int64(b.(int8)))
	case uint8:
		return cmpUint64(uint64(a.(uint8)), uint64(b.(uint8)))
	case int16:
		return cmpInt64(int64(a.(int16)), int64(b.(int16)))
	case uint16:
		return cmpUint64(uint64(a.(uint16)), uint64(b.(uint16)))
	case int32:
		return cmpInt64(int64(a.(int32)), int64(b.(int32)))
	case uint32:
		return cmpUint64(uint64(a.(uint32)), uint64(b.(uint32)))
	case int64:
		return cmpInt64(a.(int64), b.(int64))
	case uint64:
		return cmpUint64(a.(uint64), b.(uint64))
	case uintptr:
		return cmpUint64(uint64(a.(uintptr)), uint64(b.(uintptr)))
	}

	return 0
}

func cmpInt64(a, b int64) int {
	if a < b {
		return -1
	}
	return 1
}

func cmpUint64(a, b uint64) int {
	if a < b {
		return -1
	}
	return 1
}

func cmpFloat32(a, b float32) int {
	if a < b {
		return -1
	}
	return 1
}

func cmpFloat64(a, b float64) int {
	if a < b {
		return -1
	}
	return 1
}

func cmpComplex64(a, b complex64) int {
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}

func cmpComplex128(a, b complex128) int {
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}

//Reverse returns a comparator reverse to cmp
func Reverse(cmp Comparator) Comparator {
	return func(a, b interface{}) int {
		return -cmp(a, b)
	}
}

// IntComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func IntComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int) < b.(int) {
		return -1
	}
	return 1
}

// UintComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func UintComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint) < b.(uint) {
		return -1
	}
	return 1
}

// Int8Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Int8Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int8) < b.(int8) {
		return -1
	}
	return 1
}

// Uint8Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Uint8Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint8) < b.(uint8) {
		return -1
	}
	return 1
}

// Int16Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Int16Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int16) < b.(int16) {
		return -1
	}
	return 1
}

// Uint16Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Uint16Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint16) < b.(uint16) {
		return -1
	}
	return 1
}

// Int32Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Int32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int32) < b.(int32) {
		return -1
	}
	return 1
}

// Uint32Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Uint32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint32) < b.(uint32) {
		return -1
	}
	return 1
}

// Int64Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Int64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int64) < b.(int64) {
		return -1
	}
	return 1
}

// Uint64Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Uint64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint64) < b.(uint64) {
		return -1
	}
	return 1
}

// Float32Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Float32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float32) < b.(float32) {
		return -1
	}
	return 1
}

// Float64Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Float64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float64) < b.(float64) {
		return -1
	}
	return 1
}

// StringComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func StringComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(string) < b.(string) {
		return -1
	}
	return 1
}

// UintptrComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func UintptrComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uintptr) < b.(uintptr) {
		return -1
	}
	return 1
}

// BoolComparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func BoolComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(bool) == false && b.(bool) == true {
		return -1
	}
	return 1
}

// Complex64Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Complex64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	comA := a.(complex64)
	comB := b.(complex64)
	if real(comA) < real(comB) {
		return -1
	}
	if real(comA) == real(comB) && imag(comA) < imag(comB) {
		return -1
	}
	return 1
}

// Complex128Comparator compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
func Complex128Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	comA := a.(complex128)
	comB := b.(complex128)
	if real(comA) < real(comB) {
		return -1
	}
	if real(comA) == real(comB) && imag(comA) < imag(comB) {
		return -1
	}
	return 1
}
