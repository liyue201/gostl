package bitmap

// Bitmap is a mapping from some domain (for example, a range of integers) to bits. It is also called a bit array or bitmap index
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

// NewFromData creates a bitmap from exported data
func NewFromData(data []byte) *Bitmap {
	bitmap := &Bitmap{
		size: uint64(len(data)) * 8,
		data: data,
	}
	return bitmap
}

// Set set 1 at position
func (b *Bitmap) Set(position uint64) bool {
	if position >= b.size {
		return false
	}
	b.data[position>>3] |= 1 << (position & 0x07)
	return true
}

// Unset sets 0 at position
func (b *Bitmap) Unset(position uint64) bool {
	if position >= b.size {
		return false
	}
	b.data[position>>3] &= ^(1 << (position & 0x07))
	return true
}

// IsSet returns whether the position is set 1
func (b *Bitmap) IsSet(position uint64) bool {
	if position >= b.size {
		return false
	}
	if b.data[position>>3]&(1<<(position&0x07)) > 0 {
		return true
	}
	return false
}

// Resize resize the bitmap
func (b *Bitmap) Resize(size uint64) {
	size = (size + 7) / 8 * 8
	if b.size == size {
		return
	}
	data := make([]byte, size/8, size/8)
	copy(data, b.data)
	b.data = data
	b.size = size
}

// Size returns the bitmap's size in bit
func (b *Bitmap) Size() uint64 {
	return b.size
}

// Clear clear the bitmap's data
func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size/8, b.size/8)
}

// Data returns the internal data
func (b *Bitmap) Data() []byte {
	return b.data
}
