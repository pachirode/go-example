package main

import (
	"container/ring"
	"fmt"
)

func link() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// 内部有一个循环，不断的获取下一个的值
	r.Do(func(v any) {
		fmt.Println(v)
	})
}
