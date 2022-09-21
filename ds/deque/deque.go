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

// Deque is double-ended queue supports efficient data insertion from the head and tail, random access and iterator access.
type Deque[T any] struct {
	pool  *Pool[T]
	segs  []*Segment[T]
	begin int
	end   int
	size  int
}

// New creates a new deque
func New[T any]() *Deque[T] {
	dq := &Deque[T]{
		pool: newPool[T](),
		segs: make([]*Segment[T], 0),
	}
	return dq
}

// Size returns the amount of values in the deque
func (d *Deque[T]) Size() int {
	return d.size
}

// Empty returns true if the deque is empty,otherwise returns false.
func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) segUsed() int {
	if d.size == 0 {
		return 0
	}
	if d.end > d.begin {
		return d.end - d.begin
	}
	return len(d.segs) - d.begin + d.end
}

// PushFront pushed a value to the front of the deque
func (d *Deque[T]) PushFront(value T) {
	d.firstAvailableSegment().pushFront(value)
	d.size++
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
}

// PushBack pushed a value to the back of deque
func (d *Deque[T]) PushBack(value T) {
	d.lastAvailableSegment().pushBack(value)
	d.size++
	if d.segUsed() >= len(d.segs) {
		d.expand()
	}
}

// Insert inserts a value to the position pos of the deque
func (d *Deque[T]) Insert(pos int, value T) {
	if pos < 0 || pos > d.size {
		return
	}
	if pos == 0 {
		d.PushFront(value)
		return
	}
	if pos == d.size {
		d.PushBack(value)
		return
	}
	seg, pos := d.pos(pos)
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

func (d *Deque[T]) moveFrontInsert(seg, pos int, value T) {
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

func (d *Deque[T]) moveBackInsert(seg, pos int, value T) {
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

// Front returns the value at the first position of the deque
func (d *Deque[T]) Front() T {
	return d.firstSegment().front()
}

// Back returns the value at the last position of the deque
func (d *Deque[T]) Back() T {
	return d.lastSegment().back()
}

// At returns the value at position pos of the deque
func (d *Deque[T]) At(pos int) T {
	if pos < 0 || pos >= d.Size() {
		panic("out off range")
	}
	seg, pos := d.pos(pos)
	return d.segmentAt(seg).at(pos)
}

// Set sets the value of the deque's position pos with value val
func (d *Deque[T]) Set(pos int, val T) error {
	if pos < 0 || pos >= d.size {
		return ErrOutOffRange
	}
	seg, pos := d.pos(pos)
	d.segmentAt(seg).set(pos, val)
	return nil
}

// PopFront returns the value at the first position of the deque and removes it
func (d *Deque[T]) PopFront() T {
	if d.size == 0 {
		panic("deque is empty")
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

// PopBack returns the value at the lase position of the deque and removes it
func (d *Deque[T]) PopBack() T {
	if d.size == 0 {
		panic("deque is empty")
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

// EraseAt erases the element at the position pos
func (d *Deque[T]) EraseAt(pos int) {
	if pos < 0 || pos >= d.size {
		return
	}
	seg, pos := d.pos(pos)
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

// EraseRange erases elements in range [firstPos, lastPos)
func (d *Deque[T]) EraseRange(firstPos, lastPos int) {
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

// Clear erases all elements in the deque
func (d *Deque[T]) Clear() {
	d.EraseRange(0, d.size)
}

func (d *Deque[T]) putToPool(s *Segment[T]) {
	s.clear()
	d.pool.put(s)

	if d.pool.size()*6/5 > d.segUsed() {
		d.pool.shrinkToSize(d.segUsed() / 5)
	}
}

func (d *Deque[T]) firstAvailableSegment() *Segment[T] {
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

func (d *Deque[T]) lastAvailableSegment() *Segment[T] {
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

func (d *Deque[T]) firstSegment() *Segment[T] {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.begin]
}

func (d *Deque[T]) lastSegment() *Segment[T] {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.preIndex(d.end)]
}

func (d *Deque[T]) segmentAt(seg int) *Segment[T] {
	return d.segs[(seg+d.begin)%cap(d.segs)]
}

func (d *Deque[T]) pos(position int) (int, int) {
	if position <= d.firstSegment().size()-1 {
		return 0, position
	}
	position -= d.firstSegment().size()
	return position/SegmentCapacity + 1, position % SegmentCapacity
}

func (d *Deque[T]) expand() {
	newCapacity := d.segUsed() * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	seg := make([]*Segment[T], newCapacity)
	for i := 0; i < d.segUsed(); i++ {
		seg[i] = d.segs[(d.begin+i)%d.segUsed()]
	}
	u := d.segUsed()

	d.begin = 0
	d.end = u
	d.segs = seg
}

//shrinkIfNeeded shrinks the deque if it has too many unused space.
func (d *Deque[T]) shrinkIfNeeded() {
	if int(float64(d.segUsed()*2)*1.2) < cap(d.segs) {
		newCapacity := cap(d.segs) / 2
		seg := make([]*Segment[T], newCapacity)
		for i := 0; i < d.segUsed(); i++ {
			seg[i] = d.segs[(d.begin+i)%cap(d.segs)]
		}
		u := d.segUsed()
		d.begin = 0
		d.end = u
		d.segs = seg
	}
}

func (d *Deque[T]) nextIndex(index int) int {
	return (index + 1) % cap(d.segs)
}

func (d *Deque[T]) preIndex(index int) int {
	return (index - 1 + cap(d.segs)) % cap(d.segs)
}

// String returns a string representation of the deque
func (d *Deque[T]) String() string {
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

// Begin returns an iterator of the deque with the first position
func (d *Deque[T]) Begin() *DequeIterator[T] {
	return d.First()
}

// End returns an iterator of the deque with the position d.Size()
func (d *Deque[T]) End() *DequeIterator[T] {
	return d.IterAt(d.Size())
}

// First returns an iterator of the deque with the first position
func (d *Deque[T]) First() *DequeIterator[T] {
	return d.IterAt(0)
}

// Last returns an iterator of the deque with the last position
func (d *Deque[T]) Last() *DequeIterator[T] {
	return d.IterAt(d.Size() - 1)
}

// IterAt returns an iterator of the deque with the position pos
func (d *Deque[T]) IterAt(pos int) *DequeIterator[T] {
	return &DequeIterator[T]{dq: d, position: pos}
}
