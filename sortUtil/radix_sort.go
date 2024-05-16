package sortUtil

// 这里只实现了整数类型的基数排序

// func (d *Data[T]) countSort()

// 获取数字的位数
func getDigit(num, d int) int {
	for i := 1; i < d; i++ {
		num /= 10
	}
	return num % 10
}

// 计数排序
func countingSort(arr []int, d int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		count[getDigit(arr[i], d)]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[getDigit(arr[i], d)]-1] = arr[i]
		count[getDigit(arr[i], d)]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// 基数排序
func radixSort(arr []int) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}
