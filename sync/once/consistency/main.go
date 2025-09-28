package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var i = 10
	onceBody := func() {
		i *= 2
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	// 20
	fmt.Println("i", i)
}
