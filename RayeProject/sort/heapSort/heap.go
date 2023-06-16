package heapSort

// 堆排序

func HeapSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	heapSize := len(arr)
	heapSize--
	swap(arr, 0, heapSize)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		swap(arr, 0, heapSize)
	}
}
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}
func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largest := 0
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
