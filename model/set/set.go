package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](l ...int) Set[T] {
	var m map[T]struct{}
	if len(l) != 0 {
		m = make(map[T]struct{}, l[0])
	} else {
		m = make(map[T]struct{})
	}
	return m
}
func (d Set[T]) Store(val T) {
	d[val] = struct{}{}
}
func (d Set[T]) Exists(val T) bool {
	_, ok := d[val]
	return ok
}
func (d Set[T]) Remove(val T) {
	delete(d, val)
}
func (d Set[T]) ListAll() []T {
	li := make([]T, 0, len(d))
	for key, _ := range d {
		li = append(li, key)
	}
	return li
}
