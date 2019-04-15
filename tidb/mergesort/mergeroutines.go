package main

import (
	"runtime"
	"sync"
)

type mergeParam struct {
	src []int64
	dest []int64
	low int
	mid int
	high int
	wg *sync.WaitGroup
}

func setMp(mp *mergeParam, src []int64, dest []int64, low int,
	mid int, high int, wg *sync.WaitGroup) {
	mp.src = src
	mp.dest = dest
	mp.low = low
	mp.mid = mid
	mp.high = high
	mp.wg = wg
}

var queueLen = runtime.NumCPU()

var fromCh = make(chan *mergeParam, queueLen)
var toCh = make(chan *mergeParam, queueLen)

func prepareCh() {
	fillToCh()
	for i := 0; i < queueLen; i++{
		go func() {
			for mp := range fromCh{
				merge(mp.src, mp.dest, mp.low, mp.mid, mp.high)
				mp.wg.Done()
				// 归还参数对象
				toCh <- mp
			}
		}()
	}
}

func fillToCh() {
	for len(toCh) < queueLen{
		toCh <- new(mergeParam)
	}
}


