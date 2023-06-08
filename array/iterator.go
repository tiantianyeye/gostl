package array

//array iterator
type Iterator[T any] interface {
	IsValid() bool
	Prev() Iterator[T]
	Next() Iterator[T]
	Value() T
	SetValue(value T)
	IteratorAt(position int) Iterator[T]
	Position() int
	Size() int
}

type ArrayIterator[T any] struct {
	array    *Array[T]
	position int
}

func (iter *ArrayIterator[T]) IsValid() bool {
	if iter.array == nil || iter.position < 0 || iter.position >= iter.array.Size() {
		return false
	}
	return true
}

func (iter *ArrayIterator[T]) Prev() Iterator[T] {
	iter.position--
	return iter
}

func (iter *ArrayIterator[T]) Next() Iterator[T] {
	iter.position++
	return iter
}

func (iter *ArrayIterator[T]) Value() T {
	return iter.array.At(iter.position)
}

func (iter *ArrayIterator[T]) SetValue(value T) {
	iter.array.Set(iter.position, value)
}

func (iter *ArrayIterator[T]) IteratorAt(position int) Iterator[T] {
	iter.position = position
	return iter
}

func (iter *ArrayIterator[T]) Position() int {
	return iter.position
}

func (iter *ArrayIterator[T]) Size() int {
	return iter.array.Size()
}
