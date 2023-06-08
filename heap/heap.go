package heap

type Less[T any] interface {
	Less(*T) bool
}

type IHeap[T Less[T]] interface {
	Pop() *T
	Push(*T)
	At(int) *T
	Len() int
	Fix(int)
	Remove(int)
	IsEmpty() bool
}

type Heap[T Less[T]] struct {
	values []*T
}

func NewHeap[T Less[T]](size int) IHeap[T] {
	return &Heap[T]{values: make([]*T, 0, size)}
}

func (h *Heap[T]) Pop() *T {
	if h.Len() <= 0 {
		panic("heap is empty")
	}

	last := h.Len() - 1
	h.swap(0, last)
	ret := h.values[last]
	h.values = h.values[:last]
	h.down(0)
	return ret
}

func (h *Heap[T]) Push(val *T) {
	h.values = append(h.values, val)
	last := h.Len() - 1
	h.up(last)
}

func (h *Heap[T]) Fix(idx int) {
	if idx < 0 || idx >= h.Len() {
		panic("fix err idx out off range")
	}
	if !h.up(idx) {
		h.down(idx)
	}
}

func (h *Heap[T]) Remove(idx int) {
	if h.Len() <= 0 {
		panic("heap is empty")
	}

	last := h.Len() - 1
	if idx == last {
		h.values = h.values[:last]
	} else {
		h.swap(idx, last)
		h.values = h.values[:last]
		if !h.up(idx) {
			h.down(idx)
		}
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Len() <= 0
}

func (h *Heap[T]) At(idx int) *T {
	if idx < 0 || idx >= h.Len() {
		panic("at err idx out off range")
	}
	return h.values[idx]
}

func (h *Heap[T]) Len() int {
	if h.values == nil {
		return 0
	}
	return len(h.values)
}

func (h *Heap[T]) up(idx int) bool {
	curIdx := idx
	for curIdx > 0 {
		upIdx := (curIdx - 1) / 2
		curVal := h.values[curIdx]
		upVal := h.values[upIdx]

		if (*curVal).Less(upVal) {
			h.swap(curIdx, upIdx)
			curIdx = upIdx
			continue
		}
		break
	}
	return curIdx < idx
}

func (h *Heap[T]) down(idx int) bool {
	hLen := h.Len() - 1
	curIdx := idx
	for curIdx < hLen {
		downIdx := curIdx*2 + 1
		if downIdx > hLen {
			break
		}

		curVal := h.values[curIdx]

		if downIdx+1 <= hLen {
			downVal := h.values[downIdx]
			downNext := h.values[downIdx+1]
			if (*downNext).Less(downVal) {
				downIdx++
			}
		}

		downVal := h.values[downIdx]
		if (*downVal).Less(curVal) {
			h.swap(curIdx, downIdx)
			curIdx = downIdx
			continue
		}
		break
	}

	return curIdx > idx
}

func (h *Heap[T]) swap(a, b int) {
	h.values[a], h.values[b] = h.values[b], h.values[a]
}
