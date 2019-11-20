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
func (this Slice) Len() int {
	return len(this)
}

func (this Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val
}

func (this Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

///////////////////////////////////////////
// IntSlice functions
func (this IntSlice) Len() int {
	return len(this)
}

func (this IntSlice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this IntSlice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(int)
}

func (this IntSlice) Begin() *SliceIterator {
	return this.First()
}

func (this IntSlice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this IntSlice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this IntSlice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// UIntSlice functions
func (this UIntSlice) Len() int {
	return len(this)
}

func (this UIntSlice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this UIntSlice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(uint)
}

func (this UIntSlice) Begin() *SliceIterator {
	return this.First()
}

func (this UIntSlice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this UIntSlice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this UIntSlice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// Int8Slice functions
func (this Int8Slice) Len() int {
	return len(this)
}

func (this Int8Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Int8Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(int8)
}

func (this Int8Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Int8Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Int8Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Int8Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// UInt8Slice functions
func (this UInt8Slice) Len() int {
	return len(this)
}

func (this UInt8Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this UInt8Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(uint8)
}

func (this UInt8Slice) Begin() *SliceIterator {
	return this.First()
}

func (this UInt8Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this UInt8Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this UInt8Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// Int32Slice functions
func (this Int32Slice) Len() int {
	return len(this)
}

func (this Int32Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Int32Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(int32)
}

func (this Int32Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Int32Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

func (this Int32Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Int32Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// UInt32Slice functions
func (this UInt32Slice) Len() int {
	return len(this)
}

func (this UInt32Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this UInt32Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(uint32)
}

func (this UInt32Slice) Begin() *SliceIterator {
	return this.First()
}

func (this UInt32Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this UInt32Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this UInt32Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// Int64Slice functions
func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Int64Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(int64)
}

func (this Int64Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Int64Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Int64Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Int64Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// UInt64Slice functions
func (this UInt64Slice) Len() int {
	return len(this)
}

func (this UInt64Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this UInt64Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(uint64)
}

func (this UInt64Slice) Begin() *SliceIterator {
	return this.First()
}

func (this UInt64Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this UInt64Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this UInt64Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// Float32Slice functions
func (this Float32Slice) Len() int {
	return len(this)
}

func (this Float32Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Float32Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(float32)
}

func (this Float32Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Float32Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Float32Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Float32Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// Float64Slice functions
func (this Float64Slice) Len() int {
	return len(this)
}

func (this Float64Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Float64Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(float64)
}

func (this Float64Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Float64Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Float64Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Float64Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}

///////////////////////////////////////
// StringSlice functions
func (this StringSlice) Len() int {
	return len(this)
}

func (this StringSlice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this StringSlice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val.(string)
}

func (this StringSlice) Begin() *SliceIterator {
	return this.First()
}

func (this StringSlice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this StringSlice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this StringSlice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len() - 1,
	}
}
