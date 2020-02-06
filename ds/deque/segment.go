package deque

//Segment is a fixed capacity ring
type Segment struct {
	data  []interface{}
	begin int
	end   int
	nSize int
}

func newSegment(capacity int) *Segment {
	return &Segment{
		data: make([]interface{}, capacity),
	}
}

func (s *Segment) pushBack(value interface{}) {
	s.data[s.end] = value
	s.end = s.nextIndex(s.end)
	s.nSize++
}

func (s *Segment) pushFront(val interface{}) {
	s.begin = s.preIndex(s.begin)
	s.data[s.begin] = val
	s.nSize++
}

func (s *Segment) insert(position int, value interface{}) {
	if position < s.nSize-position {
		//move the front pos items
		idx := s.preIndex(s.begin)
		for i := 0; i < position; i++ {
			s.data[idx] = s.data[s.nextIndex(idx)]
			idx = s.nextIndex(idx)
		}
		s.data[idx] = value
		s.begin = s.preIndex(s.begin)
	} else {
		//move the back pos items
		idx := s.end
		for i := 0; i < s.nSize-position; i++ {
			s.data[idx] = s.data[s.preIndex(idx)]
			idx = s.preIndex(idx)
		}
		s.data[idx] = value
		s.end = s.nextIndex(s.end)
	}
	s.nSize++
}

func (s *Segment) popBack() interface{} {
	s.end = s.preIndex(s.end)
	val := s.data[s.end]
	s.data[s.end] = nil
	s.nSize--
	return val
}

func (s *Segment) popFront() interface{} {
	val := s.data[s.begin]
	s.data[s.begin] = nil
	s.begin = s.nextIndex(s.begin)
	s.nSize--
	return val
}

func (s *Segment) eraseAt(position int) {
	if position < s.nSize-position {
		for i := position; i > 0; i-- {
			index := (i + s.begin) % s.capacity()
			preIndex := (i - 1 + s.begin) % s.capacity()
			s.data[index] = s.data[preIndex]
		}
		s.data[s.begin] = nil
		s.begin = s.nextIndex(s.begin)
	} else {
		for i := position; i < s.nSize; i++ {
			index := (i + s.begin) % s.capacity()
			nextIndex := (i + 1 + s.begin) % s.capacity()
			s.data[index] = s.data[nextIndex]
		}
		s.data[s.preIndex(s.end)] = nil
		s.end = s.preIndex(s.end)
	}
	s.nSize--
}

func (s *Segment) size() int {
	return s.nSize
}

func (s *Segment) capacity() int {
	return len(s.data)
}

func (s *Segment) full() bool {
	return s.nSize == len(s.data)
}

func (s *Segment) empty() bool {
	return s.nSize == 0
}

func (s *Segment) nextIndex(index int) int {
	return (index + 1) % s.capacity()
}

func (s *Segment) preIndex(index int) int {
	return (index - 1 + s.capacity()) % s.capacity()
}

func (s *Segment) at(position int) interface{} {
	if position < 0 || position >= s.nSize {
		return nil
	}
	return s.data[(position+s.begin)%s.capacity()]
}

func (s *Segment) set(position int, val interface{}) {
	if position < 0 || position >= len(s.data) {
		return
	}
	s.data[(position+s.begin)%s.capacity()] = val
}

func (s *Segment) back() interface{} {
	return s.at(s.nSize - 1)
}

func (s *Segment) front() interface{} {
	return s.at(0)
}

func (s *Segment) clear() {
	if s.nSize > 0 {
		for i := s.begin; i != s.end; i = (i + 1) % len(s.data) {
			s.data[i] = nil
		}
	}
	s.begin = 0
	s.end = 0
	s.nSize = 0
}
