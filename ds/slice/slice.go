package slice

// Interface of Slice for iterator
type ISlice interface {
	Len() int
	At(position int) interface{}
	Set(position int, val interface{})
}

// Slice definition
type Slice []interface{}
type IntSlice []int
type UIntSlice []uint
type Int8Slice []int8
type UInt8Slice []uint8
type Int32Slice []int32
type UInt32Slice []uint32
type Int64Slice []int64
type UInt64Slice []uint64
type Float32Slice []float32
type Float64Slice []float64
type StringSlice []string

// Tells the compiler XXSlice implements ISlice
var _ ISlice = Slice(nil)
var _ ISlice = IntSlice(nil)
var _ ISlice = UIntSlice(nil)
var _ ISlice = Int8Slice(nil)
var _ ISlice = UInt8Slice(nil)
var _ ISlice = Int32Slice(nil)
var _ ISlice = Int32Slice(nil)
var _ ISlice = Int64Slice(nil)
var _ ISlice = Float32Slice(nil)
var _ ISlice = Float64Slice(nil)
var _ ISlice = StringSlice(nil)

///////////////////////////////////////////////////
// Slice functions
func (s Slice) Len() int {
	return len(s)
}

func (s Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val
}

func (s Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

///////////////////////////////////////////
// IntSlice functions
func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s IntSlice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(int)
}

func (s IntSlice) Begin() *SliceIterator {
	return s.First()
}

func (s IntSlice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s IntSlice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s IntSlice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// UIntSlice functions
func (s UIntSlice) Len() int {
	return len(s)
}

func (s UIntSlice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s UIntSlice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(uint)
}

func (s UIntSlice) Begin() *SliceIterator {
	return s.First()
}

func (s UIntSlice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s UIntSlice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s UIntSlice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// Int8Slice functions
func (s Int8Slice) Len() int {
	return len(s)
}

func (s Int8Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Int8Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(int8)
}

func (s Int8Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Int8Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s Int8Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Int8Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// UInt8Slice functions
func (s UInt8Slice) Len() int {
	return len(s)
}

func (s UInt8Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s UInt8Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(uint8)
}

func (s UInt8Slice) Begin() *SliceIterator {
	return s.First()
}

func (s UInt8Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s UInt8Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s UInt8Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// Int32Slice functions
func (s Int32Slice) Len() int {
	return len(s)
}

func (s Int32Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Int32Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(int32)
}

func (s Int32Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Int32Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

func (s Int32Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Int32Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// UInt32Slice functions
func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s UInt32Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(uint32)
}

func (s UInt32Slice) Begin() *SliceIterator {
	return s.First()
}

func (s UInt32Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s UInt32Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s UInt32Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// Int64Slice functions
func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Int64Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(int64)
}

func (s Int64Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Int64Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s Int64Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Int64Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// UInt64Slice functions
func (s UInt64Slice) Len() int {
	return len(s)
}

func (s UInt64Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s UInt64Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(uint64)
}

func (s UInt64Slice) Begin() *SliceIterator {
	return s.First()
}

func (s UInt64Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s UInt64Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s UInt64Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// Float32Slice functions
func (s Float32Slice) Len() int {
	return len(s)
}

func (s Float32Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Float32Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(float32)
}

func (s Float32Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Float32Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s Float32Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Float32Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// Float64Slice functions
func (s Float64Slice) Len() int {
	return len(s)
}

func (s Float64Slice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s Float64Slice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(float64)
}

func (s Float64Slice) Begin() *SliceIterator {
	return s.First()
}

func (s Float64Slice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s Float64Slice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s Float64Slice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}

///////////////////////////////////////
// StringSlice functions
func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) At(position int) interface{} {
	if position < 0 || position > s.Len() {
		return nil
	}
	return s[position]
}

func (s StringSlice) Set(position int, val interface{}) {
	if position < 0 || position > s.Len() {
		return
	}
	s[position] = val.(string)
}

func (s StringSlice) Begin() *SliceIterator {
	return s.First()
}

func (s StringSlice) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

func (s StringSlice) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

func (s StringSlice) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len() - 1,
	}
}
