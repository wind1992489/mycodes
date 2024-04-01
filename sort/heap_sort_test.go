package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {
	numCount := 100
	s := make([]int, numCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numCount; i++ {
		s[i] = r.Intn(1000)
	}
	fmt.Println(s)
	Heap_sort(s)
	fmt.Println(s)
}
