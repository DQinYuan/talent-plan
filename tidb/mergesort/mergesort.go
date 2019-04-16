/*
非递归版归并排序的并行化
 */
package main

import (
	"sync"
)

const max = 1 << 11

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	dst := make([]int64, len(src))
	wg := new(sync.WaitGroup)

	if len(src) <= max{
		mUntil(1, src, dst, 0, len(src), wg, false)
		return
	}

	waitNum := len(src) / max
	wg.Add(waitNum)
	i := 0
	for i <= len(src) - max{
		// 先将max大小的使用单线程归并分别排序好
		go func(low int) {
			defer wg.Done()
			mUntil(1, src, dst, low, low + max, wg, false)
		}(i)
		i += max
	}

	if i < len(src){
		mUntil(1, src, dst, i, len(src), wg, false)
	}

	wg.Wait()

	mUntil(max, src, dst, 0, len(src),wg, true)
}

// 从from长度的块开始归并排序src, par表示是否并行, 排序的范围为从low(包括)到high(不包括), par表示是否并发
func mUntil(from int, src []int64, dst []int64, low int, high int, wg *sync.WaitGroup, par bool){
	var temp []int64

	i := from
	counter := 0

	for i < high - low{
		mSort(src, dst, i, low, high, wg, par)
		i *= 2
		temp = dst
		dst = src
		src = temp
		counter += 1
	}

	if counter % 2 == 0{
		return
	}

	copy(dst[low:high], src[low:high])
}

// 以k为归并块的粒度对数组从low(包括)到high(不包括)进行整体归并
func mSort(src []int64, dest []int64, k int, low int, high int, wg *sync.WaitGroup, par bool) {
	i := low
	len := high - low

	if par{
		waitNum := len / (2 * k)
		wg.Add(waitNum)

		for i <= high - 2 * k{
			go func(i int) {
				defer wg.Done()
				merge(src, dest, i, i + k, i + 2 * k)
			}(i)
			i += 2 * k
		}

	} else {

		for i <= high - 2 * k{
			merge(src, dest, i, i + k, i + 2 * k)
			i += 2 * k
		}

	}


	if i < high - k {
		merge(src, dest, i, i + k, high)
	} else {
		copy(dest[i:high], src[i:high])
	}

	if par == true{
		wg.Wait()
	}
}

// 将src的low(包括)到mid(不包括)块和mid(包括)到high(不包括)块归并到dest
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
