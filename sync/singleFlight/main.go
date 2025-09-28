package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	cache        = make(map[string]string)
	mu           sync.RWMutex
	requestGroup singleflight.Group
)

func GetValue(key string) string {
	mu.RLock()
	val, ok := cache[key]
	mu.RUnlock()
	if !ok {

		fmt.Printf("%v not in cache\n", key)

		result, _, _ := requestGroup.Do(key, func() (interface{}, error) {
			val = key
			mu.Lock()
			cache[key] = val
			mu.Unlock()

			return val, nil
		})

		return result.(string)
	}
	return val
}

func main() {
	var wg sync.WaitGroup
	keys := []string{"1", "2", "3", "4"}

	for _, key := range keys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetValue(key)
			fmt.Printf("Result: %v : %v\n", key, GetValue(key))
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("===================================")

	for _, key := range keys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Result2: %v : %v\n", key, GetValue(key))
		}()
	}
}
