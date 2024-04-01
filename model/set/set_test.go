package set

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	d := NewSet[int](10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		d.Store(r.Intn(10))
	}
	fmt.Println(d.ListAll())
}
