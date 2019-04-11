package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	dst := make([]int64, len(src))
	copy(dst, src)
	mSort(dst, src, 0, len(src))
}

//从src归并到dest
func mSort(src []int64, dest []int64, low int, high int) {
	if high-low <= 1 {
		return
	}

	mid := (low + high) >> 1
	mSort(dest, src, low, mid)
	mSort(dest, src, mid, high)

	for i, p, q := low, low, mid; i < high; i++ {
		if q >= high || (p < mid && src[p] < src[q]) {
			dest[i] = src[p]
			p++
		} else {
			dest[i] = src[q]
			q++
		}
	}
}
