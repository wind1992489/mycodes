package sortUtil

func (d *Data[T]) HeapSort() {
	d.buildMaxHeap()
	for i := len(d.List) - 1; i > 0; i-- {
		d.Swap(0, i)
		d.heapfy(0, i-1)
	}
}
func (d *Data[T]) buildMaxHeap() {
	lastIdx := len(d.List) - 1
	for i := parent(lastIdx); i >= 0; i-- {
		l := left(i)
		r := right(i)
		largest := i
		if l <= lastIdx && d.At(l) > d.At(largest) {
			largest = l
		}
		if r <= lastIdx && d.At(r) > d.At(largest) {
			largest = r
		}
		if largest != i {
			d.Swap(i, largest)
			d.heapfy(largest, lastIdx)
		}
	}
}
func (d *Data[T]) heapfy(start int, end int) {
	l := left(start)
	r := right(start)
	largest := start
	if l <= end && d.At(l) > d.At(largest) {
		largest = l
	}
	if r <= end && d.At(r) > d.At(largest) {
		largest = r
	}
	if largest != start {
		d.Swap(start, largest)
		d.heapfy(largest, end)
	}
}

func left(i int) int {
	return i*2 + 1
}
func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}
