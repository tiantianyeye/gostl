package array

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	array := New[int](100)
	iter := array.Begin()

	idx := 1
	for iter.IsValid() {
		iter.SetValue(idx)
		idx++
		iter.Next()
	}

	iter = iter.IteratorAt(iter.Size() - 1)
	for iter.IsValid() {
		fmt.Println("a ...any:", iter.Value())
		iter.Prev()
	}
}

func TestBit(t *testing.T) {

	pos := 8
	t1 := pos >> 3
	t2 := 1 << (pos & 0x07)
	fmt.Printf("test bit i:%d, t111111---------:%d, t2222222--------------:%d \n", pos, t1, t2)
}
