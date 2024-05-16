package sortUtil

import "testing"

func TestNewSortUtil(t *testing.T) {
	d := NewSortUtil[float64](100, 1000)
	d.ListAll()
	d.QuickSort()
	d.ListAll()
	d.Shuffle()
	d.ListAll()
	d.HeapSort()
	d.ListAll()
	d.Shuffle()
	d.ListAll()
	d.MergeSort()
	d.ListAll()
}
