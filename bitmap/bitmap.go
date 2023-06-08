package bitmap

type Bitmap struct {
	data []byte
	size uint64
}

func New(size uint64) *Bitmap {
	size = (size + 7) / 8 * 8
	return &Bitmap{
		data: make([]byte, size/8, size/8),
		size: size,
	}
}

func NewFromData(data []byte) *Bitmap {
	return &Bitmap{
		data: data,
		size: uint64(len(data)) * 8,
	}
}

func (b *Bitmap) Set(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] |= 1 << (pos & 0x07)
	return true
}

func (b *Bitmap) Unset(pos uint64) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] &= ^(1 << (pos & 0x07))
	return true
}

func (b *Bitmap) Isset(pos uint64) bool {
	if pos >= b.size {
		return false
	}

	if (b.data[pos>>3] & (1 << (pos & 0x07))) > 0 {
		return true
	}
	return false
}

func (b *Bitmap) Size() uint64 {
	return b.size
}

func (b *Bitmap) Resize(size uint64) {
	size = (size + 7) / 8 * 8
	if size == b.size {
		return
	}
	data := make([]byte, size/8, size/8)
	copy(data, b.data)
	b.data = data
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size/8, b.size/8)
}

func (b *Bitmap) Data() []byte {
	return b.data
}
