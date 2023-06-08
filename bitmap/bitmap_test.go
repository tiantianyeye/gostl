package bitmap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPos(t *testing.T) {
	bm := New(100)
	for i := 0; i < 10; i++ {
		t := rand.Intn(100)
		bm.Set(uint64(t))

		fmt.Printf("set pos:%d status:%v \n", t, bm.Isset(uint64(t)))
	}

	for i := 0; i < 100; i++ {
		if bm.Isset(uint64(i)) {
			fmt.Printf("get pod:%d, status:%v \n", i, bm.Isset(uint64(i)))
		}
	}

}
