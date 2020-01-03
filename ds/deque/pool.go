package deque

type Pool struct {
	segs []*Segment
}

func NewPool() *Pool {
	return &Pool{segs: make([]*Segment, 0)}
}

func (p *Pool) Get() *Segment {
	if len(p.segs) == 0 {
		return NewSegment(SegmentCapacity)
	}
	s := p.segs[len(p.segs)-1]
	p.segs = p.segs[len(p.segs)-1:]
	return s
}

func (p *Pool) Put(s *Segment) {
	p.segs = append(p.segs, s)
}

func (p *Pool) ShrinkToSize(size int) {
	if len(p.segs) > size {
		newSeg := make([]*Segment, size)
		copy(newSeg, p.segs)
		p.segs = newSeg
	}
}

func (p *Pool) Size() int {
	return len(p.segs)
}
