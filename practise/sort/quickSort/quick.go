package quickSort

//快排

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	p := arr[right]
	i, j := left, right
	for i < j {
		for i < j && arr[i] <= p {
			i++
		}
		arr[j] = arr[i]
		for i < j && arr[j] >= p {
			j--
		}
		arr[i] = arr[j]
	}
	arr[i] = p
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
