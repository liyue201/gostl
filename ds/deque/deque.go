package deque

import (
	"errors"
	"fmt"
)

const (
	SegmentCapacity = 64
)

var ErrOutOffRange = errors.New("out off range")

//Deque is a ring
type Deque struct {
	pool    *Pool
	segs    [] *Segment
	begin   int
	end     int
	segUsed int
	size    int
}

func New() *Deque {
	dq := &Deque{
		pool:    NewPool(),
		segs:    make([]*Segment, 1),
		end:     1,
		segUsed: 1,
	}
	dq.segs[0] = dq.pool.Get()
	return dq
}

//Size returns the size of deque
func (this *Deque) Size() int {
	return this.size
}

//Empty returns true if the Deque is empty,otherwise returns false.
func (this *Deque) Empty() bool {
	return this.size == 0
}

func (this *Deque) PushFront(value interface{}) {
	this.firstAvailableSegment().PushFront(value)
	this.size++
}

func (this *Deque) PushBack(value interface{}) {
	this.lastAvailableSegment().PushBack(value)
	this.size++
}

func (this *Deque) Insert(position int, value interface{}) {
	if position < 0 || position > this.size {
		return
	}
	if position == 0 {
		this.PushFront(value)
		return
	}
	if position == this.size {
		this.PushBack(value)
		return
	}
	seg, pos := this.pos(position)
	if seg < this.segUsed-seg {
		if this.firstSegment().Full() {
			if this.segUsed == cap(this.segs) {
				this.expand()
			}
			if pos == 0 {
				pos = SegmentCapacity - 1
			} else {
				seg++
				pos--
			}
			this.begin = this.preIndex(this.begin)
			this.segs[this.begin] = this.pool.Get()
			this.segUsed++
		}
		for i := 0; i < seg; i++ {
			cur := this.segmentAt(i)
			next := this.segmentAt(i + 1)
			cur.PushBack(next.PopFront())
		}
		this.segmentAt(seg).Insert(pos, value)
	} else {
		// move back
		if this.lastSegment().Full() {
			if this.segUsed == len(this.segs) {
				this.expand()
			}
			this.segs[this.end] = this.pool.Get()
			this.end = this.nextIndex(this.end)
			this.segUsed++
		}
		for i := this.segUsed - 1; i > seg; i-- {
			cur := this.segmentAt(i)
			pre := this.segmentAt(i - 1)
			cur.PushFront(pre.PopBack())
		}
		this.segmentAt(seg).Insert(pos, value)
	}
	this.size++
}

func (this *Deque) Front() interface{} {
	return this.firstSegment().Front()
}

func (this *Deque) Back() interface{} {
	return this.lastSegment().Back()
}

func (this *Deque) At(position int) interface{} {
	if position < 0 || position >= this.Size() {
		return nil
	}
	seg, pos := this.pos(position)
	return this.segmentAt(seg).At(pos)
}

func (this *Deque) Set(position int, val interface{}) error {
	if position < 0 || position >= this.size {
		return ErrOutOffRange
	}
	seg, pos := this.pos(position)
	this.segmentAt(seg).Set(pos, val)
	return nil
}

func (this *Deque) PopFront() interface{} {
	if this.size == 0 {
		return nil
	}
	s := this.segs[this.begin]
	if s.Size() == 1 {
		this.putToPool(this.segs[this.begin])
		this.segs[this.begin] = nil
		this.begin = this.nextIndex(this.begin)
	}
	this.size--
	this.shrinkIfNeeded()
	return s.PopFront()
}

func (this *Deque) PopBack() interface{} {
	if this.size == 0 {
		return nil
	}
	seg := this.preIndex(this.end)
	s := this.segs[seg]
	if s.Size() == 1 {
		this.putToPool(this.segs[seg])
		this.segs[seg] = nil
		this.end = seg
	}
	this.size--
	this.shrinkIfNeeded()
	return s.PopBack()
}

func (this *Deque) EraseAt(position int) {
	if position < 0 || position >= this.size {
		return
	}
	seg, pos := this.pos(position)
	this.segmentAt(seg).EraseAt(pos)
	if seg < this.size-seg-1 {
		for i := seg; i > 0; i-- {
			cur := this.segmentAt(i)
			pre := this.segmentAt(i - 1)
			cur.PushFront(pre.PopBack())
		}
		if this.firstSegment().Empty() {
			this.putToPool(this.firstSegment())
			this.segs[this.begin] = nil
			this.begin = this.nextIndex(this.begin)
			this.segUsed--
			this.shrinkIfNeeded()
		}
	} else {
		for i := seg; i < this.segUsed-1; i++ {
			cur := this.segmentAt(i)
			next := this.segmentAt(i + 1)
			cur.PushBack(next.PopFront())
		}
		if this.lastSegment().Empty() {
			this.putToPool(this.lastSegment())
			this.segs[this.preIndex(this.end)] = nil
			this.end = this.preIndex(this.end)
			this.segUsed--
			this.shrinkIfNeeded()
		}
	}
	this.size--
}

// EraseRange erases data in range [firstPos, lastPos)
func (this *Deque) EraseRange(firstPos, lastPos int) {
	if firstPos < 0 || firstPos >= lastPos || lastPos > this.size {
		return
	}
	s1, p1 := this.pos(firstPos)
	s2, p2 := this.pos(lastPos - 1)
	segDelFrom := s1 + 1
	segDelTo := s2 - 1
	if p1 == 0 {
		segDelFrom = s1
	}
	if p2 == SegmentCapacity-1 {
		segDelTo = s2
	}
	num := segDelTo - segDelFrom + 1

	if this.segUsed-segDelFrom < segDelTo {
		// move back
		pos := segDelFrom
		count := 0
		for ; pos < this.segUsed-num; pos++ {
			index := (this.begin + pos) % len(this.segs)
			nextIndex := (this.begin + pos + num) % len(this.segs)
			if count < num {
				this.segs[index].Clear()
				this.putToPool(this.segs[index])
			}
			count++
			this.segs[index] = this.segs[nextIndex]
		}

		for ; pos < this.segUsed; pos++ {
			index := (this.begin + pos) % len(this.segs)
			this.segs[index] = nil
		}
		this.end = (this.end - num + len(this.segs)) % len(this.segs)

	} else {
		// move front
		pos := segDelTo
		count := 0
		for ; pos >= num; pos-- {
			index := (this.begin + pos) % len(this.segs)
			preIndex := (this.begin + pos - num) % len(this.segs)
			if count < num {
				this.segs[index].Clear()
				this.putToPool(this.segs[index])
			}
			count++

			this.segs[index] = this.segs[preIndex]
		}
		for ; pos >= 0; pos-- {
			index := (this.begin + pos) % len(this.segs)
			this.segs[index] = nil
		}
		this.begin = (this.begin + num) % len(this.segs)
	}
	this.segUsed -= num
	this.size -= num * SegmentCapacity

	left := lastPos - firstPos - num*SegmentCapacity
	erasePos := firstPos
	for ; left > 0; left-- {
		this.EraseAt(erasePos)
	}
}

func (this *Deque) putToPool(s *Segment) {
	this.pool.Put(s)
	if this.pool.Size()*6/5 > this.segUsed {
		this.pool.ShrinkToSize(this.segUsed / 5)
	}
}

func (this *Deque) firstAvailableSegment() *Segment {
	if !this.firstSegment().Full() {
		return this.firstSegment()
	}
	if this.segUsed == len(this.segs) {
		this.expand()
	}
	this.begin = this.preIndex(this.begin)
	s := this.pool.Get()
	this.segs[this.begin] = s
	this.segUsed++
	return s
}

func (this *Deque) lastAvailableSegment() *Segment {
	if !this.lastSegment().Full() {
		return this.lastSegment()
	}
	if this.segUsed == len(this.segs) {
		this.expand()
	}
	s := this.pool.Get()
	this.segs[this.end] = s
	this.end = this.nextIndex(this.end)
	this.segUsed++
	return s
}

func (this *Deque) firstSegment() *Segment {
	return this.segs[this.begin]
}

func (this *Deque) lastSegment() *Segment {
	return this.segs[this.preIndex(this.end)]
}

func (this *Deque) segmentAt(seg int) *Segment {
	return this.segs[(seg+this.begin)%cap(this.segs)]
}

//pos returns the segment number and position in segment
func (this *Deque) pos(position int) (int, int) {
	if position <= this.firstSegment().Size()-1 {
		return 0, position
	}
	position -= this.firstSegment().Size()
	return position/SegmentCapacity + 1, position % SegmentCapacity
}

func (this *Deque) expand() {
	newCapacity := this.segUsed * 2
	seg := make([]*Segment, newCapacity)
	for i := 0; i < this.segUsed; i++ {
		seg[i] = this.segs[(this.begin+i)%this.segUsed]
	}
	this.segs = seg
	this.begin = 0
	this.end = this.segUsed
}

//shrinkIfNeeded shrinks the Deque if is has too many unused space .
func (this *Deque) shrinkIfNeeded() {
	if int(float64(this.segUsed*2)*1.2) < cap(this.segs) {
		newCapacity := cap(this.segs) / 2
		seg := make([]*Segment, newCapacity)
		for i := 0; i < this.segUsed; i++ {
			seg[i] = this.segs[(this.begin+i)%cap(this.segs)]
		}
		this.segs = seg
		this.begin = 0
		this.end = this.segUsed
	}
}

func (this *Deque) nextIndex(index int) int {
	return (index + 1) % cap(this.segs)
}

func (this *Deque) preIndex(index int) int {
	return (index - 1 + cap(this.segs)) % cap(this.segs)
}

func (this *Deque) String() string {
	str := "["
	for i := 0; i < this.Size(); i++ {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", this.At(i))
	}
	str += "]"

	return str
}

///////////////////////////////////////////////////
//iterator API
func (this *Deque) Begin() *DequeIterator {
	return this.First()
}

func (this *Deque) End() *DequeIterator {
	return this.IterAt(this.Size())
}

func (this *Deque) First() *DequeIterator {
	return this.IterAt(0)
}

func (this *Deque) Last() *DequeIterator {
	return this.IterAt(this.Size() - 1)
}

func (this *Deque) IterAt(position int) *DequeIterator {
	return &DequeIterator{dq: this, position: position}
}
