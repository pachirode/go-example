package main

import (
	"fmt"
	"io"
	"os"
)

func demo() {
	fmt.Println("1")

	defer func() {
		fmt.Println("2")
		defer fmt.Println("3")
		fmt.Println("4")
	}()

	fmt.Println("5")
}

func noting() int {
	r := 2
	defer func() {
		fmt.Println("r: ", r)
		r *= 3
	}()

	// 2
	return r
}

func write() (r int) {
	defer func() {
		fmt.Println("r: ", r)
		r *= 3
	}()

	return 2
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	//demo()
	fmt.Println(noting())
	fmt.Println(write())
	CopyFile("tmp", "go.mod")
}
