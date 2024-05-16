package sort

import (
	"cmp"
)

// func Quick_sort(s []int, start, end int) {
// 	if start >= end {
// 		return
// 	}
// 	idx := find_split(s, start, end)
// 	Quick_sort(s, start, idx-1)
// 	Quick_sort(s, idx+1, end)
// }
// func find_split(s []int, start, end int) int {
// 	q := s[start]
// 	p := start
// 	for i := start; i <= end; i++ {
// 		if s[i] < q {
// 			s[i], s[p+1] = s[p+1], s[i]
// 			p++
// 		}
// 	}
// 	s[p], s[start] = s[start], s[p]
// 	return p
// }

func QuickSort[T cmp.Ordered](s []T) {
	if len(s) <= 1 {
		return
	}
	idx := split(s)
	QuickSort(s[:idx])
	QuickSort(s[idx+1:])
}

func split[T cmp.Ordered](s []T) int {
	idx := 0
	q := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < q {
			idx++
			s[idx], s[i] = s[i], s[idx]
		}
	}
	s[idx], s[0] = s[0], s[idx]
	return idx
}
