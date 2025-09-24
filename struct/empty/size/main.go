package main

import (
	"fmt"
	"unsafe"
)

type Empty struct {
}

func printAddress() {
	var s1 struct{}
	fmt.Printf("addr: %p, size: %d\n", &s1, unsafe.Sizeof(s1))
}

func judgeAddressValues() {
	var (
		a struct{}
		b struct{}
		c struct{}
		d struct{}
	)

	println("&a:", &a)
	println("&b:", &b)
	println("&c:", &c)
	println("&d:", &d)

	println("&a == &b:", &a == &b)
	x := &a
	y := &b
	println("x == y:", x == y)

	fmt.Printf("&c(%p) == &d(%p): %t\n", &c, &d, &c == &d)

	//Output:
	//&a: 0xc000053eff
	//&b: 0xc000053eff
	//&c: 0x8b9680
	//&d: 0x8b9680
	//&a == &b: false
	//x == y: false

}

func addressSizeByPosition() {
	type A struct {
		x struct{}
		y int
		z string
		k int64
	}

	type B struct {
		x int
		z struct{}
		y string
		k int64
	}

	type C struct {
		x int
		y string
		z struct{}
		k int64
	}

	type D struct {
		x int
		y string
		z int64
		k struct{}
	}

	a := A{}
	b := B{}
	c := C{}
	d := D{}
	fmt.Printf("struct a size: %d\n", unsafe.Sizeof(a))
	fmt.Printf("struct b size: %d\n", unsafe.Sizeof(b))
	fmt.Printf("struct c size: %d\n", unsafe.Sizeof(c))
	fmt.Printf("struct d size: %d\n", unsafe.Sizeof(d))
	// Output:
	//struct a size: 32
	//struct b size: 24
	//struct c size: 24

}

func main() {
	addressSizeByPosition()
}
