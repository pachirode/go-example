package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	var g errgroup.Group
	g.SetLimit(3)

	for i := 1; i <= 10; i++ {
		if g.TryGo(func() error {
			fmt.Printf("Goroutine %d is starting\n", i)
			time.Sleep(2 * time.Second) // 模拟工作
			fmt.Printf("Goroutine %d is done\n", i)
			return nil
		}) {
			fmt.Printf("Goroutine %d started successfully\n", i)
		} else {
			fmt.Printf("Goroutine %d could not start (limit reached)\n", i)
		}
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered an error: %v\n", err)
	}

	fmt.Println("All goroutines complete.")
}
