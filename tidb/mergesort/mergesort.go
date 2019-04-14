package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	dst := make([]int64, len(src))
	i := 1
	counter := 0
	var temp []int64
	for i < len(src){
		mSort(src, dst, i)
		i *= 2
		temp = dst
		dst = src
		src = temp
		counter += 1
	}

	if counter % 2 == 0{
		return
	}

	copy(dst, src)
}


//从src归并到dest
func mSort(src []int64, dest []int64, k int) {
	i := 0
	for i <= len(src) - 2 * k{
		merge(src, dest, i, i + k, i + 2 * k)
		i += 2 * k
	}

	if i < len(src) - k {
		merge(src, dest, i, i + k, len(src))
	} else {
		copy(dest[i:], src[i:])
	}
}

func merge(src []int64, dest []int64, low int, mid int, high int)  {
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
