package main

import (
	"fmt"
	"time"
)

func f() {
	defer fmt.Println("defer 1")
	fmt.Println(1)
	panic("woah")
	defer fmt.Println("defer 2")
	fmt.Println(2)
}

func g() {
	fmt.Println("calling g")
	// 子 goroutine 中发生 panic，主 goroutine 也会退出
	go f()
	time.Sleep(2 * time.Second)
	fmt.Println("called g")
}

func main() {
	//f()

	g()
}
