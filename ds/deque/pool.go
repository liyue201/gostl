package deque

type Pool struct {
	segs []*Segment
}

func NewPool() *Pool {
	return &Pool{segs: make([]*Segment, 0)}
}

func (this *Pool) Get() *Segment {
	if len(this.segs) == 0 {
		return NewSegment(SegmentCapacity)
	}
	s := this.segs[len(this.segs)-1]
	this.segs = this.segs[len(this.segs)-1:]
	return s
}

func (this *Pool) Put(s *Segment) {
	this.segs = append(this.segs, s)
}

func (this *Pool) ShrinkToSize(size int) {
	if len(this.segs) > size {
		newSeg := make([]*Segment, size)
		copy(newSeg, this.segs)
		this.segs = newSeg
	}
}

func (this *Pool) Size() int {
	return len(this.segs)
}
