package slice

type Slice []interface{}

func (this Slice) Len() int {
	return len(this)
}

func (this Slice) At(position int) interface{} {
	if position < 0 || position > this.Len() {
		return nil
	}
	return this[position]
}

func (this Slice) Set(position int, val interface{}) {
	if position < 0 || position > this.Len() {
		return
	}
	this[position] = val
}

func (this Slice) Begin() *SliceIterator {
	return this.First()
}

func (this Slice) End() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}

func (this Slice) First() *SliceIterator {
	return &SliceIterator{s: this,
		position: 0,
	}
}

func (this Slice) Last() *SliceIterator {
	return &SliceIterator{s: this,
		position: this.Len(),
	}
}
