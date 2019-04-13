package main

import (
	"sort"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		MergeSort(src)
	}
}

func BenchmarkNormalSort(b *testing.B) {
	// 16M个整数
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	}
}

type sorted []int64

func (s sorted) Len() int  {
	return len([]int64(s))
}

func (s sorted) Less(i, j int) bool {
	return i < j
}

func (s sorted) Swap(i, j int){
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func BenchmarkSortSort(b *testing.B) {
	// 16M个整数
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		sort.Sort(sorted(src))
	}
}


func BenchmarkQuickSort(b *testing.B) {
	// 16M个整数
	numElements := 16 << 20
	src := make([]int64, numElements)
	original := make([]int64, numElements)
	prepare(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, original)
		b.StartTimer()
		quickSort(src)
	}
}
