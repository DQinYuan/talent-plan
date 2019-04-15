package main

import (
	"github.com/pingcap/check"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var _ = check.Suite(&sortTestSuite{})

//测试*sortTestSuite中所有Test开头的方法
func TestQ(t *testing.T) {
	check.TestingT(t)
}

func prepareQ(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int63()
	}
}

type sortTestSuiteQ struct{}

func (s *sortTestSuite) TestQuickSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		quickSort(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}