package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	numCount := 10
	s := make([]int, numCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numCount; i++ {
		s[i] = r.Intn(1000)
	}
	fmt.Println(s)
	MergeSort(s)
	fmt.Println(s)
}
