package array

import "fmt"

// array is a fixed slice
type Array[T any] struct {
	values []T
}

func New[T any](size int) *Array[T] {
	return &Array[T]{values: make([]T, size, size)}
}

func NewFromSlice[T any](array []T) *Array[T] {
	return &Array[T]{values: array}
}

func (a *Array[T]) Size() int {
	return len(a.values)
}

func (a *Array[T]) Empty() bool {
	if a.values == nil || len(a.values) <= 0 {
		return true
	}
	return false
}

func (a *Array[T]) Data() []T {
	return a.values
}

func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.values)
}

func (a *Array[T]) At(pos int) T {
	if pos < 0 || pos >= len(a.values) {
		panic("index out of range")
	}
	return a.values[pos]
}

func (a *Array[T]) Set(pos int, val T) {
	if pos < 0 || pos >= len(a.values) {
		panic("index out of range")
	}
	a.values[pos] = val
}

func (a *Array[T]) Swap(pos1, pos2 int) {
	if pos1 < 0 || pos1 >= len(a.values) || pos2 < 0 || pos2 >= len(a.values) {
		panic("index out of range")
	}

	a.values[pos1], a.values[pos2] = a.values[pos2], a.values[pos1]
}

func (a *Array[T]) Begin() Iterator[T] {
	return &ArrayIterator[T]{array: a, position: 0}
}
