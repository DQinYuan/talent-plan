package main

func quickSort(arr []int64) {
	qSort(arr, 0, len(arr) - 1)
}

func qSort(arr []int64, low int, high int) {
	if low < high{
		pivot := partition(arr, low, high)
		qSort(arr, low, pivot - 1)
		qSort(arr, pivot + 1, high)
	}
}

func partition(arr []int64, low int, high int) int {
	pivotKey := arr[low]
	for low < high{
		for low < high && arr[high] >= pivotKey{
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivotKey{
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivotKey

	return low
}
