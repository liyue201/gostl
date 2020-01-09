package deque

import (
	"errors"
	"fmt"
)

// Constants definition
const (
	SegmentCapacity = 128
)

// Define internal errors
var (
	ErrOutOffRange = errors.New("out off range")
)

// Deque supports efficient data insertion from the head and tail, random access and iterator access.
type Deque struct {
	pool  *Pool
	segs  []*Segment
	begin int
	end   int
	size  int
}

// New news a deque
func New() *Deque {
	dq := &Deque{
		pool: newPool(),
		segs: make([]*Segment, 0),
	}
	return dq
}

// Size returns the size of deque
func (d *Deque) Size() int {
	return d.size
}

// Empty returns true if the Deque is empty,otherwise returns false.
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

// PushFront pushed value to the front of d
func (d *Deque) PushFront(value interface{}) {
	d.firstAvailableSegment().pushFront(value)
	d.size++
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
}

// PushBack pushed value to the back of d
func (d *Deque) PushBack(value interface{}) {
	d.lastAvailableSegment().pushBack(value)
	d.size++
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
}

// Insert inserts value to the position of d
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
		// seg is closer to the front
		d.moveFrontInsert(seg, pos, value)
	} else {
		// seg is closer to the back
		d.moveBackInsert(seg, pos, value)
	}
	d.size++
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
}

func (d *Deque) moveFrontInsert(seg, pos int, value interface{}) {
	if d.firstSegment().full() {
		if d.segUsed() >= len(d.segs) {
			d.expand()
		}
		d.begin = d.preIndex(d.begin)
		d.segs[d.begin] = d.pool.get()
		if pos == 0 {
			pos = SegmentCapacity - 1
		} else {
			seg++
			pos--
		}
	} else {
		if pos == 0 {
			seg--
			pos = SegmentCapacity - 1
		} else {
			if seg != 0 {
				pos--
			}
		}
	}
	for i := 0; i < seg; i++ {
		cur := d.segmentAt(i)
		next := d.segmentAt(i + 1)
		cur.pushBack(next.popFront())
	}
	d.segmentAt(seg).insert(pos, value)
}

func (d *Deque) moveBackInsert(seg, pos int, value interface{}) {
	// move back
	if d.lastSegment().full() {
		if d.segUsed() >= len(d.segs) {
			d.expand()
		}
		d.segs[d.end] = d.pool.get()
		d.end = d.nextIndex(d.end)
	}
	for i := d.segUsed() - 1; i > seg; i-- {
		cur := d.segmentAt(i)
		pre := d.segmentAt(i - 1)
		cur.pushFront(pre.popBack())
	}
	d.segmentAt(seg).insert(pos, value)
}

// Front returns the front value of d
func (d *Deque) Front() interface{} {
	return d.firstSegment().front()
}

// Back returns the back value of d
func (d *Deque) Back() interface{} {
	return d.lastSegment().back()
}

// At returns the value of d at position
func (d *Deque) At(position int) interface{} {
	if position < 0 || position >= d.Size() {
		return nil
	}
	seg, pos := d.pos(position)
	return d.segmentAt(seg).at(pos)
}

// Set sets the value of d at position
func (d *Deque) Set(position int, val interface{}) error {
	if position < 0 || position >= d.size {
		return ErrOutOffRange
	}
	seg, pos := d.pos(position)
	d.segmentAt(seg).set(pos, val)
	return nil
}

// PopFront returns the font value fo d, and removes it
func (d *Deque) PopFront() interface{} {
	if d.size == 0 {
		return nil
	}
	s := d.segs[d.begin]
	v := s.popFront()
	if s.size() == 0 {
		d.putToPool(s)
		d.segs[d.begin] = nil
		d.begin = d.nextIndex(d.begin)
	}
	d.size--
	d.shrinkIfNeeded()
	return v
}

// PopBack returns the back value fo d, and removes it
func (d *Deque) PopBack() interface{} {
	if d.size == 0 {
		return nil
	}
	seg := d.preIndex(d.end)
	s := d.segs[seg]
	v := s.popBack()

	if s.size() == 0 {
		d.putToPool(s)
		d.segs[seg] = nil
		d.end = seg
	}

	d.size--
	d.shrinkIfNeeded()
	return v
}

// EraseAt erases the item at position
func (d *Deque) EraseAt(position int) {
	if position < 0 || position >= d.size {
		return
	}
	seg, pos := d.pos(position)
	d.segmentAt(seg).eraseAt(pos)
	if seg < d.size-seg-1 {
		for i := seg; i > 0; i-- {
			cur := d.segmentAt(i)
			pre := d.segmentAt(i - 1)
			cur.pushFront(pre.popBack())
		}
		if d.firstSegment().empty() {
			d.putToPool(d.firstSegment())
			d.segs[d.begin] = nil
			d.begin = d.nextIndex(d.begin)
			d.shrinkIfNeeded()
		}
	} else {
		for i := seg; i < d.segUsed()-1; i++ {
			cur := d.segmentAt(i)
			next := d.segmentAt(i + 1)
			cur.pushBack(next.popFront())
		}
		if d.lastSegment().empty() {
			d.putToPool(d.lastSegment())
			d.segs[d.preIndex(d.end)] = nil
			d.end = d.preIndex(d.end)
			d.shrinkIfNeeded()
		}
	}
	d.size--
}

// EraseRange erases items in range [firstPos, lastPos)
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
	s.clear()
	d.pool.put(s)

	if d.pool.size()*6/5 > d.segUsed() {
		d.pool.shrinkToSize(d.segUsed() / 5)
	}
}

func (d *Deque) firstAvailableSegment() *Segment {
	if d.firstSegment() != nil && !d.firstSegment().full() {
		return d.firstSegment()
	}
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
	if d.firstSegment() == nil || d.firstSegment().full() {
		d.begin = d.preIndex(d.begin)
		s := d.pool.get()
		d.segs[d.begin] = s
		return s
	}
	return d.firstSegment()
}

func (d *Deque) lastAvailableSegment() *Segment {
	if d.lastSegment() != nil && !d.lastSegment().full() {
		return d.lastSegment()
	}
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
	if d.lastSegment() == nil || d.lastSegment().full() {
		s := d.pool.get()
		d.segs[d.end] = s
		d.end = d.nextIndex(d.end)
		return s
	}
	return d.lastSegment()
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
	if position <= d.firstSegment().size()-1 {
		return 0, position
	}
	position -= d.firstSegment().size()
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
	u := d.segUsed()

	d.begin = 0
	d.end = u
	d.segs = seg
}

//shrinkIfNeeded shrinks the Deque if is has too many unused space .
func (d *Deque) shrinkIfNeeded() {
	if int(float64(d.segUsed()*2)*1.2) < cap(d.segs) {
		newCapacity := cap(d.segs) / 2
		seg := make([]*Segment, newCapacity)
		for i := 0; i < d.segUsed(); i++ {
			seg[i] = d.segs[(d.begin+i)%cap(d.segs)]
		}
		u := d.segUsed()
		d.begin = 0
		d.end = u
		d.segs = seg
	}
}

func (d *Deque) nextIndex(index int) int {
	return (index + 1) % cap(d.segs)
}

func (d *Deque) preIndex(index int) int {
	return (index - 1 + cap(d.segs)) % cap(d.segs)
}

// String returns d in string format
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

// Begin returns the first iterator of d
func (d *Deque) Begin() *DequeIterator {
	return d.First()
}

// End returns the end iterator of d
func (d *Deque) End() *DequeIterator {
	return d.IterAt(d.Size())
}

// First returns the first iterator of d
func (d *Deque) First() *DequeIterator {
	return d.IterAt(0)
}

// Last returns the last iterator of d
func (d *Deque) Last() *DequeIterator {
	return d.IterAt(d.Size() - 1)
}

// IterAt returns the iterator at position of d
func (d *Deque) IterAt(position int) *DequeIterator {
	return &DequeIterator{dq: d, position: position}
}
