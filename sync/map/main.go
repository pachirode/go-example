package main

import (
	"fmt"
	"sync"
)

func base() {
	var sMap sync.Map

	sMap.Store("name", "test")
	sMap.Store("age", 1)
	sMap.Store("location", "this")

	if value, ok := sMap.Load("name"); ok {
		println("name", value)
	}

	sMap.Delete("name")

	sMap.Range(func(key, value any) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true
	})
}

// fatal error: concurrent map read and map write
func baseGoroutine() {
	m := make(map[string]int)

	go func() {
		for {
			m["k"] = 1
			println("set k: ", 1)
		}
	}()

	go func() {
		for {
			v, _ := m["k"]
			println("get k: ", v)
		}
	}()

	select {}
}

func mutexGoroutine() {
	m := NewRwMap()

	go func() {
		for {
			m.Set("k", 1)
			println("set k: ", 1)
		}
	}()

	go func() {
		for {
			v, _ := m.Get("k")
			println("get k: ", v)
		}
	}()

	select {}
}

func main() {
	//base()
	//baseGoroutine()
	//mutexGoroutine()
}
