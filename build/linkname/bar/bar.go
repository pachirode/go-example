package bar

import (
	_ "unsafe"
)

func test() {
	println("bar.test")
}

// Push

//go:linkname testPush github.com/pachirode/go-example/build/linkname/foo.TestPush
func testPush() {
	println("bar.testPush")
}

// Handshake

//go:linkname testHandshake
func testHandshake() {
	println("test.handshake")
}
