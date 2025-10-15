package main

import (
	"fmt"
	"time"

	"github.com/gammazero/workerpool"
)

func base() {
	wp := workerpool.New(2)
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, r := range cases {
		wp.Submit(func() {
			fmt.Println(time.Now(), r)
			time.Sleep(1 * time.Second)
		})
	}

	wp.StopWait()
}

func main() {
	base()
}
