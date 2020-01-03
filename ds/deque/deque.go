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
	pool  *Pool
	segs  []*Segment
	begin int
	end   int
	size  int
}

func New() *Deque {
	dq := &Deque{
		pool: NewPool(),
		segs: make([]*Segment, 0),
	}
	return dq
}

//Size returns the size of deque
func (d *Deque) Size() int {
	return d.size
}

//Empty returns true if the Deque is empty,otherwise returns false.
func (d *Deque) Empty() bool {
	return d.size == 0
}

func (d *Deque) segUsed() int {
	if d.size == 0 {
		return 0
	}
	if d.end > d.begin {
		return d.end - d.begin
	}
	return len(d.segs) - d.begin + d.end
}

func (d *Deque) PushFront(value interface{}) {
	d.firstAvailableSegment().PushFront(value)
	d.size++
}

func (d *Deque) PushBack(value interface{}) {
	d.lastAvailableSegment().PushBack(value)
	d.size++
}

func (d *Deque) Insert(position int, value interface{}) {
	if position < 0 || position > d.size {
		return
	}
	if position == 0 {
		d.PushFront(value)
		return
	}
	if position == d.size {
		d.PushBack(value)
		return
	}
	seg, pos := d.pos(position)
	if seg < d.segUsed()-seg {
		if d.firstSegment().Full() {
			if d.segUsed() == cap(d.segs) {
				d.expand()
			}
			if pos == 0 {
				pos = SegmentCapacity - 1
			} else {
				seg++
				pos--
			}
			d.begin = d.preIndex(d.begin)
			d.segs[d.begin] = d.pool.Get()
		}
		for i := 0; i < seg; i++ {
			cur := d.segmentAt(i)
			next := d.segmentAt(i + 1)
			cur.PushBack(next.PopFront())
		}
		d.segmentAt(seg).Insert(pos, value)
	} else {
		// move back
		if d.lastSegment().Full() {
			if d.segUsed() == len(d.segs) {
				d.expand()
			}
			d.segs[d.end] = d.pool.Get()
			d.end = d.nextIndex(d.end)
		}
		for i := d.segUsed() - 1; i > seg; i-- {
			cur := d.segmentAt(i)
			pre := d.segmentAt(i - 1)
			cur.PushFront(pre.PopBack())
		}
		d.segmentAt(seg).Insert(pos, value)
	}
	d.size++
}

func (d *Deque) Front() interface{} {
	return d.firstSegment().Front()
}

func (d *Deque) Back() interface{} {
	return d.lastSegment().Back()
}

func (d *Deque) At(position int) interface{} {
	if position < 0 || position >= d.Size() {
		return nil
	}
	seg, pos := d.pos(position)
	return d.segmentAt(seg).At(pos)
}

func (d *Deque) Set(position int, val interface{}) error {
	if position < 0 || position >= d.size {
		return ErrOutOffRange
	}
	seg, pos := d.pos(position)
	d.segmentAt(seg).Set(pos, val)
	return nil
}

func (d *Deque) PopFront() interface{} {
	if d.size == 0 {
		return nil
	}
	s := d.segs[d.begin]
	if s.Size() == 1 {
		d.putToPool(d.segs[d.begin])
		d.segs[d.begin] = nil
		d.begin = d.nextIndex(d.begin)
	}
	d.size--
	d.shrinkIfNeeded()
	return s.PopFront()
}

func (d *Deque) PopBack() interface{} {
	if d.size == 0 {
		return nil
	}
	seg := d.preIndex(d.end)
	s := d.segs[seg]
	if s.Size() == 1 {
		d.putToPool(d.segs[seg])
		d.segs[seg] = nil
		d.end = seg
	}
	d.size--
	d.shrinkIfNeeded()
	return s.PopBack()
}

func (d *Deque) EraseAt(position int) {
	if position < 0 || position >= d.size {
		return
	}
	seg, pos := d.pos(position)
	d.segmentAt(seg).EraseAt(pos)
	if seg < d.size-seg-1 {
		for i := seg; i > 0; i-- {
			cur := d.segmentAt(i)
			pre := d.segmentAt(i - 1)
			cur.PushFront(pre.PopBack())
		}
		if d.firstSegment().Empty() {
			d.putToPool(d.firstSegment())
			d.segs[d.begin] = nil
			d.begin = d.nextIndex(d.begin)
			d.shrinkIfNeeded()
		}
	} else {
		for i := seg; i < d.segUsed()-1; i++ {
			cur := d.segmentAt(i)
			next := d.segmentAt(i + 1)
			cur.PushBack(next.PopFront())
		}
		if d.lastSegment().Empty() {
			d.putToPool(d.lastSegment())
			d.segs[d.preIndex(d.end)] = nil
			d.end = d.preIndex(d.end)
			d.shrinkIfNeeded()
		}
	}
	d.size--
}

// EraseRange erases data in range [firstPos, lastPos)
func (d *Deque) EraseRange(firstPos, lastPos int) {
	if firstPos < 0 || firstPos >= lastPos || lastPos > d.size {
		return
	}
	num := lastPos - firstPos
	if d.size-firstPos < lastPos {
		// move back
		for pos := firstPos; pos+num < d.size; pos++ {
			d.Set(pos, d.At(pos+num))
		}
		for ; num > 0; num-- {
			d.PopBack()
		}
	} else {
		// move front
		for pos := lastPos - 1; pos-num >= 0; pos-- {
			d.Set(pos, d.At(pos-num))
		}
		for ; num > 0; num-- {
			d.PopFront()
		}
	}
}

// Clear erases all elements in deque
func (d *Deque) Clear() {
	d.EraseRange(0, d.size)
}

func (d *Deque) putToPool(s *Segment) {
	d.pool.Put(s)
	if d.pool.Size()*6/5 > d.segUsed() {
		d.pool.ShrinkToSize(d.segUsed() / 5)
	}
}

func (d *Deque) firstAvailableSegment() *Segment {
	if d.firstSegment() != nil && !d.firstSegment().Full() {
		return d.firstSegment()
	}
	if d.segUsed() == len(d.segs) {
		d.expand()
	}
	d.begin = d.preIndex(d.begin)
	s := d.pool.Get()
	d.segs[d.begin] = s
	return s
}

func (d *Deque) lastAvailableSegment() *Segment {
	if d.lastSegment() != nil && !d.lastSegment().Full() {
		return d.lastSegment()
	}
	if d.segUsed() == len(d.segs) {
		d.expand()
	}
	s := d.pool.Get()
	d.segs[d.end] = s
	d.end = d.nextIndex(d.end)
	return s
}

func (d *Deque) firstSegment() *Segment {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.begin]
}

func (d *Deque) lastSegment() *Segment {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.preIndex(d.end)]
}

func (d *Deque) segmentAt(seg int) *Segment {
	return d.segs[(seg+d.begin)%cap(d.segs)]
}

//pos returns the segment number and position in segment
func (d *Deque) pos(position int) (int, int) {
	if position <= d.firstSegment().Size()-1 {
		return 0, position
	}
	position -= d.firstSegment().Size()
	return position/SegmentCapacity + 1, position % SegmentCapacity
}

func (d *Deque) expand() {
	newCapacity := d.segUsed() * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	seg := make([]*Segment, newCapacity)
	for i := 0; i < d.segUsed(); i++ {
		seg[i] = d.segs[(d.begin+i)%d.segUsed()]
	}
	d.segs = seg
	d.begin = 0
	d.end = d.segUsed()
}

//shrinkIfNeeded shrinks the Deque if is has too many unused space .
func (d *Deque) shrinkIfNeeded() {
	if int(float64(d.segUsed()*2)*1.2) < cap(d.segs) {
		newCapacity := cap(d.segs) / 2
		seg := make([]*Segment, newCapacity)
		for i := 0; i < d.segUsed(); i++ {
			seg[i] = d.segs[(d.begin+i)%cap(d.segs)]
		}
		d.segs = seg
		d.begin = 0
		d.end = d.segUsed()
	}
}

func (d *Deque) nextIndex(index int) int {
	return (index + 1) % cap(d.segs)
}

func (d *Deque) preIndex(index int) int {
	return (index - 1 + cap(d.segs)) % cap(d.segs)
}

func (d *Deque) String() string {
	str := "["
	for i := 0; i < d.Size(); i++ {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", d.At(i))
	}
	str += "]"

	return str
}

///////////////////////////////////////////////////
//iterator API
func (d *Deque) Begin() *DequeIterator {
	return d.First()
}

func (d *Deque) End() *DequeIterator {
	return d.IterAt(d.Size())
}

func (d *Deque) First() *DequeIterator {
	return d.IterAt(0)
}

func (d *Deque) Last() *DequeIterator {
	return d.IterAt(d.Size() - 1)
}

func (d *Deque) IterAt(position int) *DequeIterator {
	return &DequeIterator{dq: d, position: position}
}
