package sort

import "cmp"

func Heap_sort[T cmp.Ordered, li ~[]T](s li) {
	if len(s) <= 1 {
		return
	}
	buildMaxHeap(s)
	for i := len(s) - 1; i > 0; i-- {
		swap(s, 0, i)
		heapfy(s, 0, i-1)
	}
}
func swap[T cmp.Ordered, li ~[]T](s li, i, j int) {
	if i == j {
		return
	}
	s[i], s[j] = s[j], s[i]
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
func buildMaxHeap[T cmp.Ordered, li ~[]T](s li) { // 从下到上构建
	lastID := len(s) - 1
	firstParentNodeID := (lastID - 1) / 2
	for i := firstParentNodeID; i >= 0; i-- {
		largestID := i
		leftID := left(i)
		rightID := right(i)
		if leftID <= lastID && s[leftID] > s[largestID] {
			largestID = leftID
		}
		if rightID <= lastID && s[rightID] > s[largestID] {
			largestID = rightID
		}
		swap(s, i, largestID)
		heapfy(s, largestID, lastID) // 这里由于发生了交换，需要对子节点重新heapfy
	}
}
func heapfy[T cmp.Ordered, li ~[]T](s li, startID, lastID int) { // 从上到下构建大顶堆
	largestID := startID
	leftID := left(startID)
	rightID := right(startID)
	if leftID <= lastID && s[leftID] > s[largestID] {
		largestID = leftID
	}
	if rightID <= lastID && s[rightID] > s[largestID] {
		largestID = rightID
	}
	if largestID != startID {
		swap(s, startID, largestID)
		heapfy(s, largestID, lastID)
	}
	return
}
