package main

import "fmt"

func f() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	f()
}
