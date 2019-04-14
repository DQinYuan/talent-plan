package main

import "runtime"

type mergeParam struct {
	src []int64
	dest []int64
	low int
	mid int
	high int
}

var fromCh = make(chan *mergeParam, runtime.NumCPU())
var toCh = make(chan *mergeParam, runtime.NumCPU())

func prepareCh() {
	fillToCh()
	for i := 0; i < runtime.NumCPU(); i++{
		go func() {
			for mp := range fromCh{
				merge(mp.src, mp.dest, mp.low, mp.mid, mp.high)
				// 归还参数对象
				toCh <- mp
			}
		}()
	}
}

func fillToCh() {
	for i := 0; i < runtime.NumCPU(); i++{
		toCh <- new(mergeParam)
	}
}


