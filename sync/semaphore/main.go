package main

import (
	"context"
	"log"
	"runtime"

	"golang.org/x/sync/semaphore"
)

func main() {
	var (
		ctx         = context.TODO()
		maxResource = runtime.GOMAXPROCS(0)
		sem         = semaphore.NewWeighted(int64(maxResource))
		tasks       = make([]int, 32)
	)

	for i := range tasks {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}

		go func(i int) {
			defer sem.Release(1)
			tasks[i] = task(i)
		}(i)
	}

	if err := sem.Acquire(ctx, int64(maxResource)); err != nil {
		log.Fatal("Failed to ge all resource", err)
	}
}

func task(i int) int {
	log.Println("task", i)
	return i
}
