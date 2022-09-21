package deque

// Pool is a memory pool for holding Segments
type Pool[T any] struct {
	segs []*Segment[T]
}

func newPool[T any]() *Pool[T] {
	return &Pool[T]{segs: make([]*Segment[T], 0)}
}

func (p *Pool[T]) get() *Segment[T] {
	if len(p.segs) == 0 {
		return newSegment[T](SegmentCapacity)
	}
	s := p.segs[len(p.segs)-1]
	p.segs = p.segs[:len(p.segs)-1]
	return s
}

func (p *Pool[T]) put(s *Segment[T]) {
	p.segs = append(p.segs, s)
}

func (p *Pool[T]) shrinkToSize(size int) {
	if len(p.segs) > size {
		newSeg := make([]*Segment[T], size)
		copy(newSeg, p.segs)
		p.segs = newSeg
	}
}

func (p *Pool[T]) size() int {
	return len(p.segs)
}
