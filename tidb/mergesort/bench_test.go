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
