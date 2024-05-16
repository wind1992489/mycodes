package sortUtil

import (
	"cmp"
	"fmt"
	"math/rand"
	"time"
)

type Data[T cmp.Ordered] struct {
	List      []T
	QueueSize int
	Limit     int
}

func NewSortUtil[T ~int | ~int32 | ~int64 | ~float32 | ~float64](size int, limit int) *Data[T] {
	d := Data[T]{
		QueueSize: size,
		Limit:     limit,
	}
	d.Init()
	return &d
}
func (d *Data[T]) Init() {
	size := d.QueueSize
	limit := d.Limit
	d.List = make([]T, 0, size)
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch r := any(d).(type) {
	case *Data[int]:
		for i := 0; i < size; i++ {
			r.List = append(r.List, rd.Intn(limit))
		}
	case *Data[int32]:
		for i := 0; i < size; i++ {
			r.List = append(r.List, rd.Int31n(int32(limit)))
		}
	case *Data[int64]:
		for i := 0; i < size; i++ {
			r.List = append(r.List, rd.Int63n(int64(limit)))
		}
	case *Data[float32]:
		for i := 0; i < size; i++ {
			r.List = append(r.List, rd.Float32()*float32(limit))
		}
	case *Data[float64]:
		for i := 0; i < size; i++ {
			r.List = append(r.List, rd.Float64()*float64(limit))
		}
	default:
		panic("unknown type")
	}
}
func (d *Data[T]) At(i int) T {
	if i < 0 || i >= len(d.List) {
		panic("invalid index")
	}
	return d.List[i]
}
func (d *Data[T]) Swap(i, j int) {
	if i < 0 || i >= len(d.List) {
		panic("invalid index")
	}
	if j < 0 || j >= len(d.List) {
		panic("invalid index")
	}
	d.List[i], d.List[j] = d.List[j], d.List[i]
}
func (d *Data[T]) ListAll() {
	fmt.Println(d.List)
}
func (d *Data[T]) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.List), func(i, j int) {
		d.List[i], d.List[j] = d.List[j], d.List[i]
	})
}
