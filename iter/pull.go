package main

import (
	"fmt"
	"iter"
)

func iterCase(slice []int) func(yield func(i, v int) bool) {
	return func(yield func(i int, v int) bool) {
		for i, v := range slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

func iterPullUsage() {
	s := []int{1, 2, 3, 4, 5}
	next, stop := iter.Pull2(iterCase(s))
	i, v, ok := next()
	fmt.Printf("i=%d v=%d ok=%t\n", i, v, ok)
	i, v, ok = next()
	fmt.Printf("i=%d v=%d ok=%t\n", i, v, ok)
	stop()
	i, v, ok = next()
	fmt.Printf("i=%d v=%d ok=%t\n", i, v, ok)
}
