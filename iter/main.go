package main

import "fmt"

func baseIter(yield func() bool) {
	for i := 0; i < 10; i++ {
		if !yield() {
			return
		}
	}
}

func baseIterUsage() {
	i := 0
	for range baseIter {
		// yield 函数内部执行近似逻辑
		fmt.Printf("i=%d\n", i)
		i++
		//return true
	}
}

func nIter(n int) func(yield func() bool) {
	return func(yield func() bool) {
		for i := 0; i < n; i++ {
			if !yield() {
				return
			}
		}
	}
}

func nIterUsage() {
	i := 0
	for range nIter(3) {
		fmt.Printf("i=%d\n", i)
		i++
	}
}

func judgeIter(n int) func(yield func(v int) bool) {
	return func(yield func(v int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func judgeIterUsage() {
	i := 0
	for v := range judgeIter(10) {
		if i >= 5 {
			break
		}
		fmt.Printf("%d => %d\n", i, v)
		i++
	}
}

func kvIter(slice []int) func(yield func(i, v int) bool) {
	return func(yield func(i int, v int) bool) {
		for i, v := range slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

func kvIterUsage() {
	s := []int{0, 1, 2, 3, 4}
	for i, v := range kvIter(s) {
		if i == 2 {
			continue
		}
		fmt.Printf("%d => %d\n", i, v)
	}
}

func main() {
	//baseIterUsage()
	//nIterUsage()
	//judgeIterUsage()
	//kvIterUsage()
	//gIterUsage()
	//gIter2Usage()
	iterPullUsage()
}
