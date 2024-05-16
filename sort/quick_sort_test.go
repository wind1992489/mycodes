package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuick_sort(t *testing.T) {
	numCount := 100
	s := make([]int, numCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numCount; i++ {
		s[i] = r.Intn(1000)
	}
	fmt.Println(s)
	QuickSort(s)
	fmt.Println(s)
}
