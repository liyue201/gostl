package deque

//Segment is a fixed capacity ring
type Segment[T any] struct {
	data  []T
	begin int
	end   int
	nSize int
}

func newSegment[T any](capacity int) *Segment[T] {
	return &Segment[T]{
		data: make([]T, capacity),
	}
}

func (s *Segment[T]) pushBack(value T) {
	s.data[s.end] = value
	s.end = s.nextIndex(s.end)
	s.nSize++
}

func (s *Segment[T]) pushFront(val T) {
	s.begin = s.preIndex(s.begin)
	s.data[s.begin] = val
	s.nSize++
}

func (s *Segment[T]) insert(position int, value T) {
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

func (s *Segment[T]) popBack() T {
	s.end = s.preIndex(s.end)
	val := s.data[s.end]
	//s.data[s.end] = nil
	s.nSize--
	return val
}

func (s *Segment[T]) popFront() T {
	val := s.data[s.begin]
	//s.data[s.begin] = nil
	s.begin = s.nextIndex(s.begin)
	s.nSize--
	return val
}

func (s *Segment[T]) eraseAt(position int) {
	if position < s.nSize-position {
		for i := position; i > 0; i-- {
			index := (i + s.begin) % s.capacity()
			preIndex := (i - 1 + s.begin) % s.capacity()
			s.data[index] = s.data[preIndex]
		}
		//s.data[s.begin] = nil
		s.begin = s.nextIndex(s.begin)
	} else {
		for i := position; i < s.nSize; i++ {
			index := (i + s.begin) % s.capacity()
			nextIndex := (i + 1 + s.begin) % s.capacity()
			s.data[index] = s.data[nextIndex]
		}
		//s.data[s.preIndex(s.end)] = nil
		s.end = s.preIndex(s.end)
	}
	s.nSize--
}

func (s *Segment[T]) size() int {
	return s.nSize
}

func (s *Segment[T]) capacity() int {
	return len(s.data)
}

func (s *Segment[T]) full() bool {
	return s.nSize == len(s.data)
}

func (s *Segment[T]) empty() bool {
	return s.nSize == 0
}

func (s *Segment[T]) nextIndex(index int) int {
	return (index + 1) % s.capacity()
}

func (s *Segment[T]) preIndex(index int) int {
	return (index - 1 + s.capacity()) % s.capacity()
}

func (s *Segment[T]) at(position int) T {
	if position < 0 || position >= s.nSize {
		//return nil
		panic(ErrOutOffRange.Error())
	}
	return s.data[(position+s.begin)%s.capacity()]
}

func (s *Segment[T]) set(position int, val T) {
	if position < 0 || position >= len(s.data) {
		return
	}
	s.data[(position+s.begin)%s.capacity()] = val
}

func (s *Segment[T]) back() T {
	return s.at(s.nSize - 1)
}

func (s *Segment[T]) front() T {
	return s.at(0)
}

func (s *Segment[T]) clear() {
	if s.nSize > 0 {
		for i := s.begin; i != s.end; i = (i + 1) % len(s.data) {
			//s.data[i] = nil
		}
	}
	s.begin = 0
	s.end = 0
	s.nSize = 0
}
