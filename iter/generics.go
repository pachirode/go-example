package main

import "fmt"

func gIter[Slice ~[]E, E any](s Slice) func(yield func() bool) {
	return func(yield func() bool) {
		for range s {
			if !yield() {
				return
			}
		}
	}
}

func gIterUsage() {
	s1 := []int{1, 2, 3}
	i := 0
	for range gIter(s1) {
		fmt.Printf("i=%d\n", i)
		i++
	}
}

func gIter2[Slice ~[]E, E any](s Slice) func(yield func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func gIter2Usage() {
	s1 := []int{1, 2, 3}
	for i, v := range gIter2(s1) {
		fmt.Printf("%d=%d\n", i, v)
	}
}
