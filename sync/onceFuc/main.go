package main

import (
	"sync"
)

func main() {
	onceBody := sync.OnceFunc(func() {
		panic("Only once")
	})
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			onceBody()
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
