package main

import "sync/atomic"

type S1 struct {
	a int32
	b int64
}

type S2 struct {
	a   int32
	pad uint32 // 确保内存按照 8 对齐，此处为内存填充
	b   int64
}

func main() {
	s1 := S1{}
	atomic.AddInt64(&s1.b, 1)
}
