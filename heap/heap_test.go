package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

type LessInt struct {
	val int
}

func (l LessInt) Less(less *LessInt) bool {
	return l.val < less.val
}

func TestHeap(t *testing.T) {
	heapImpl := NewHeap[LessInt](0)

	for i := int32(10); i < 100; i++ {
		num := rand.Int31n(i)
		val := &LessInt{val: int(num)}
		heapImpl.Push(val)
	}

	//heapImpl.Remove(10)

	tt := heapImpl.At(50)
	tt.val = 0
	heapImpl.Fix(50)

	ret := make([]int, 0)
	for !heapImpl.IsEmpty() {
		pop := heapImpl.Pop()
		ret = append(ret, pop.val)
	}

	fmt.Printf("-------len:%d, all:%v \n", len(ret), ret)
}
