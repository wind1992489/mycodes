package sortUtil

func (d *Data[T]) MergeSort() {
	copy(d.List, d.mergeSort(0, len(d.List)-1))
}

func (d *Data[T]) mergeSort(start, end int) []T {
	if start == end {
		return []T{d.At(start)}
	}
	mid := (start + end) / 2
	s1 := d.mergeSort(start, mid)
	s2 := d.mergeSort(mid+1, end)
	return d.merge(s1, s2)
}
func (d *Data[T]) merge(s1, s2 []T) []T {
	res := make([]T, 0, len(s1)+len(s2))
	i, j := 0, 0
loop:
	/* 这里会死循环！！！
	for ; i < len(s1) && j < len(s2) && s1[i] < s2[j]; i++ {
		res = append(res, s1[i])
	}
	for ; i < len(s1) && j < len(s2) && s1[i] > s2[j]; j++ {
		res = append(res, s2[j])
	}*/
	for ; i < len(s1) && j < len(s2) && s1[i] <= s2[j]; i++ {
		res = append(res, s1[i])
	}
	for ; i < len(s1) && j < len(s2) && s1[i] > s2[j]; j++ {
		res = append(res, s2[j])
	}
	if i < len(s1) && j < len(s2) {
		goto loop
	}
	for ; i < len(s1); i++ {
		res = append(res, s1[i])
	}
	for ; j < len(s2); j++ {
		res = append(res, s2[j])
	}
	return res
}
