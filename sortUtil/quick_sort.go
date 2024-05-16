package sortUtil

func (d *Data[T]) QuickSort() {
	if len(d.List) == 0 {
		return
	}
	d.quickSort(0, d.QueueSize-1)
}

func (d *Data[T]) quickSort(start, end int) {
	if start >= end {
		return
	}
	q := d.findSplit(start, end)
	d.quickSort(start, q-1)
	d.quickSort(q+1, end)
}

func (d *Data[T]) findSplit(start, end int) int {
	q := d.At(start)
	cursor := start
	for i := start; i <= end; i++ {
		if d.At(i) < q {
			cursor++
			d.Swap(i, cursor)
		}
	}
	d.Swap(start, cursor)
	return cursor
}
