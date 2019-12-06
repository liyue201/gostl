package bitmap

type Bitmap struct {
	data []byte
	size uint64 //bitmap's size in bit, is the multiple of 8
}

//New create a bitmap with size bit
func New(size uint64) *Bitmap {
	size = (size + 7) / 8 * 8
	bitmap := &Bitmap{
		size: size,
		data: make([]byte, size/8, size/8),
	}
	return bitmap
}

func NewFromData(data []byte) *Bitmap {
	bitmap := &Bitmap{
		size: uint64(len(data)) * 8,
		data: data,
	}
	return bitmap
}

// Set set 1 at position
func (this *Bitmap) Set(position uint64) bool {
	if position >= this.size {
		return false
	}
	this.data[position>>3] |= 1 << (position & 0x07)
	return true
}

// Unset set 0 at position
func (this *Bitmap) Unset(position uint64) bool {
	if position >= this.size {
		return false
	}
	this.data[position>>3] &= ^(1 << (position & 0x07))
	return true
}

// Unset set 0 at position
func (this *Bitmap) IsSet(position uint64) bool {
	if position >= this.size {
		return false
	}
	if this.data[position>>3]&(1<<(position&0x07)) > 0 {
		return true
	}
	return false
}

// Resize resize the bitmap
func (this *Bitmap) Resize(size uint64) {
	size = (size + 7) / 8 * 8
	if this.size == size {
		return
	}
	data := make([]byte, size/8, size/8)
	copy(data, this.data)
	this.data = data
	this.size = size
}

// Size returns the bitmap's size in bit
func (this *Bitmap) Size() uint64 {
	return this.size
}

// Clear clear the bitmap's data
func (this *Bitmap) Clear() {
	this.data = make([]byte, this.size/8, this.size/8)
}

func (this *Bitmap) Data() []byte {
	return this.data
}
