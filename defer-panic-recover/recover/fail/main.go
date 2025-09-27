package main

import "fmt"

func f() {
	m := map[int]struct{}{}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("goroutine 1", err)
			}
		}()
		for {
			m[1] = struct{}{}
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("goroutine 2", err)
			}
		}()
		for {
			m[1] = struct{}{}
		}
	}()

	select {} // 没有 case 语句，会永久阻塞
}

func main() {
	f()
}
