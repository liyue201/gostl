package deque

//Segment is a fixed size ring
type Segment struct {
	data  []interface{}
	begin int
	end   int
	size  int
}

func NewSegment(capacity int) *Segment {
	return &Segment{
		data: make([]interface{}, capacity),
	}
}

func (this *Segment) PushBack(value interface{}) {
	this.data[this.end] = value
	this.end = this.nextIndex(this.end)
	this.size++
}

func (this *Segment) PushFront(val interface{}) {
	this.begin = this.preIndex(this.begin)
	this.data[this.begin] = val
	this.size++
}

func (this *Segment) Insert(position int, value interface{}) {
	if position < this.size-position {
		//move the front pos items
		idx := this.preIndex(this.begin)
		for i := 0; i < position; i++ {
			this.data[idx] = this.data[this.nextIndex(idx)]
			idx = this.nextIndex(idx)
		}
		this.data[idx] = value
		this.begin = this.preIndex(this.begin)
	} else {
		//move the back pos items
		idx := this.end
		for i := 0; i < this.size-position; i++ {
			this.data[idx] = this.data[this.preIndex(idx)]
			idx = this.preIndex(idx)
		}
		this.data[idx] = value
		this.end = this.nextIndex(this.end)
	}
	this.size++
}

func (this *Segment) PopBack() interface{} {
	this.end = this.preIndex(this.end)
	val := this.data[this.end]
	this.data[this.end ] = nil
	this.size--
	return val
}

func (this *Segment) PopFront() interface{} {
	val := this.data[this.begin]
	this.data[this.begin] = nil
	this.begin = this.nextIndex(this.begin)
	this.size--
	return val
}

func (this *Segment) EraseAt(position int) {
	if position < this.size-position {
		for i := position; i > 0; i-- {
			index := (i + this.begin) % this.Capacity()
			preIndex := (i - 1 + this.begin) % this.Capacity()
			this.data[index] = this.data[preIndex]
		}
		this.data[this.begin] = nil
		this.begin = this.nextIndex(this.begin)
	} else {
		for i := position; i < this.size; i++ {
			index := (i + this.begin) % this.Capacity()
			nextIndex := (i + 1 + this.begin) % this.Capacity()
			this.data[index] = this.data[nextIndex]
		}
		this.data[this.preIndex(this.end)] = nil
		this.end = this.preIndex(this.end)
	}
	this.size --
}

func (this *Segment) Size() int {
	return this.size
}

func (this *Segment) Capacity() int {
	return len(this.data)
}

func (this *Segment) Full() bool {
	return this.size == len(this.data)
}

func (this *Segment) Empty() bool {
	return this.size == 0
}

func (this *Segment) nextIndex(index int) int {
	return (index + 1) % this.Capacity()
}

func (this *Segment) preIndex(index int) int {
	return (index - 1 + this.Capacity()) % this.Capacity()
}

func (this *Segment) At(position int) interface{} {
	if position < 0 || position >= this.size {
		return nil
	}
	return this.data[(position+this.begin)%this.Capacity()]
}

func (this *Segment) Set(position int, val interface{}) {
	if position < 0 || position >= len(this.data) {
		return
	}
	this.data[(position+this.begin)%this.Capacity()] = val
}

func (this *Segment) Back() interface{} {
	return this.At(this.size - 1)
}

func (this *Segment) Front() interface{} {
	return this.At(0)
}

func (this *Segment) Clear() {
	if this.size > 0 {
		for i := this.begin; i != this.end; i = (i + 1) % len(this.data) {
			this.data[i] = nil
		}
	}
	this.begin = 0
	this.end = 0
	this.size = 0
}
