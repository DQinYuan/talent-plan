package main

import "sync"


// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	prepareCh()
	dst := make([]int64, len(src))
	wg := new(sync.WaitGroup)

	mUntil(len(src), src, dst, wg)
}


func mUntil(threshold int, src []int64, dst []int64, wg *sync.WaitGroup, par bool){
	var temp []int64

	i := 1
	counter := 0

	for i < threshold{
		mSort(src, dst, i, wg, par)
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
func mSort(src []int64, dest []int64, k int, wg *sync.WaitGroup, par bool) {
	waitNum := len(src) / (2 * k)
	wg.Add(waitNum)

	i := 0
	for i <= len(src) - 2 * k{
		mp := <- toCh
		setMp(mp, src, dest, i, i + k, i + 2 * k, wg)
		fromCh <- mp
		i += 2 * k
	}

	wg.Wait()

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
