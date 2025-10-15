package main

import "runtime"

func test() {
	println(runtime.GOMAXPROCS(0))
}

func main() {
	test()
}
