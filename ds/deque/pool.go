package deque

// Pool is a memory pool for holding Segments
type Pool struct {
	segs []*Segment
}

func newPool() *Pool {
	return &Pool{segs: make([]*Segment, 0)}
}

func (p *Pool) get() *Segment {
	if len(p.segs) == 0 {
		return newSegment(SegmentCapacity)
	}
	s := p.segs[len(p.segs)-1]
	p.segs = p.segs[len(p.segs)-1:]
	return s
}

func (p *Pool) put(s *Segment) {
	p.segs = append(p.segs, s)
}

func (p *Pool) shrinkToSize(size int) {
	if len(p.segs) > size {
		newSeg := make([]*Segment, size)
		copy(newSeg, p.segs)
		p.segs = newSeg
	}
}

func (p *Pool) size() int {
	return len(p.segs)
}
