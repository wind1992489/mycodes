package sort

func heap_sort(s []int, start int, end int) {
	if start >= end {
		return
	}
	buildMaxHeap(s, start, end)
	s[start], s[end] = s[end], s[start]
	for i := end - 1; i > start; i-- {
	}
}

func buildMaxHeap(s []int, start, end int) {

}
func heapfy(s []int, start, end int) {

}
