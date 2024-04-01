package sort

import "cmp"

func MergeSort[T cmp.Ordered, li ~[]T](s li) {
	if len(s) <= 1 {
		return
	}
	idx := len(s) / 2
	MergeSort(s[:idx])
	MergeSort(s[idx:])
	s1 := Merge(s[:idx], s[idx:])
	copy(s, s1)
}

func Merge[T cmp.Ordered, li ~[]T](xs, ys li) li {
	s1 := make(li, 0, len(xs)+len(ys))
	k1, k2 := 0, 0
	for k1 < len(xs) && k2 < len(ys) {
		if xs[k1] < ys[k2] {
			s1 = append(s1, xs[k1])
			k1++
		} else {
			s1 = append(s1, ys[k2])
			k2++
		}
	}
	if k1 < len(xs) {
		s1 = append(s1, xs[k1:]...)
	}
	if k2 < len(ys) {
		s1 = append(s1, xs[k2:]...)
	}
	return s1
}
