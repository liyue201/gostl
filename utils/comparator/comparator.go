package comparator

// Should return a number:
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
type Comparator func(a, b interface{}) int

// Compare a with b
//    -1 , if a < b
//    0  , if a == b
//    1  , if a > b
// make sure a and b are both builtin type
func BuiltinTypeComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	switch a.(type) {
	case int:
		if a.(int) < b.(int) {
			return -1
		}
	case uint:
		if a.(uint) < b.(uint) {
			return -1
		}
	case int8:
		if a.(int8) < b.(int8) {
			return -1
		}
	case uint8:
		if a.(uint8) < b.(uint8) {
			return -1
		}
	case int16:
		if a.(int16) < b.(int16) {
			return -1
		}
	case uint16:
		if a.(uint16) < b.(uint16) {
			return -1
		}
	case int32:
		if a.(int32) < b.(int32) {
			return -1
		}
	case uint32:
		if a.(uint32) < b.(uint32) {
			return -1
		}
	case int64:
		if a.(int64) < b.(int64) {
			return -1
		}
	case uint64:
		if a.(uint64) < b.(uint64) {
			return -1
		}
	case float32:
		if a.(float32) < b.(float32) {
			return -1
		}
	case float64:
		if a.(float64) < b.(float64) {
			return -1
		}
	case uintptr:
		if a.(uintptr) < b.(uintptr) {
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
		comA := a.(complex64)
		comB := b.(complex64)
		if real(comA) < real(comB) {
			return -1
		}
		if real(comA) == real(comB) && imag(comA) < imag(comB) {
			return -1
		}
	case complex128:
		comA := a.(complex128)
		comB := b.(complex128)
		if real(comA) < real(comB) {
			return -1
		}
		if real(comA) == real(comB) && imag(comA) < imag(comB) {
			return -1
		}
	}
	return 1
}

//Reverse returns a comparator reverse to cmp
func Reverse(cmp Comparator) Comparator {
	return func(a, b interface{}) int {
		return -cmp(a, b)
	}
}

func IntComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int) < b.(int) {
		return -1
	}
	return 1
}

func UintComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint) < b.(uint) {
		return -1
	}
	return 1
}

func Int8Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int8) < b.(int8) {
		return -1
	}
	return 1
}

func Uint8Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint8) < b.(uint8) {
		return -1
	}
	return 1
}

func Int16Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int16) < b.(int16) {
		return -1
	}
	return 1
}

func Uint16Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint16) < b.(uint16) {
		return -1
	}
	return 1
}

func Int32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int32) < b.(int32) {
		return -1
	}
	return 1
}

func Uint32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint32) < b.(uint32) {
		return -1
	}
	return 1
}

func Int64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int64) < b.(int64) {
		return -1
	}
	return 1
}

func Uint64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint64) < b.(uint64) {
		return -1
	}
	return 1
}

func Float32Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float32) < b.(float32) {
		return -1
	}
	return 1
}

func Float64Comparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float64) < b.(float64) {
		return -1
	}
	return 1
}

func StringComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(string) < b.(string) {
		return -1
	}
	return 1
}

func UintptrComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uintptr) < b.(uintptr) {
		return -1
	}
	return 1
}

func BoolComparator(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(bool) == false && b.(bool) == true {
		return -1
	}
	return 1
}

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
