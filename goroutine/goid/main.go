package main

import (
	"fmt"
	"sync"

	"github.com/petermattis/goid"
)

func main() {
	fmt.Println("main", goid.Get())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, goid.Get())
		}()
	}
	wg.Wait()
}
