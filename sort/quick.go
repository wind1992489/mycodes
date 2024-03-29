package sort

func Quick_sort(s []int, start, end int) {
	if start >= end {
		return
	}
	idx := find_split(s, start, end)
	Quick_sort(s, start, idx-1)
	Quick_sort(s, idx+1, end)
}
func find_split(s []int, start, end int) int {
	q := s[start]
	p := start
	for i := start; i <= end; i++ {
		if s[i] < q {
			s[i], s[p+1] = s[p+1], s[i]
			p++
		}
	}
	s[p], s[start] = s[start], s[p]
	return p
}
