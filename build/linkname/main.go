package main

import (
	"github.com/pachirode/go-example/build/linkname/foo"
)

func pull() {
	foo.Test()
}

func push() {
	foo.TestPush()
}

func handshake() {
	foo.TestHandshake()
}

func main() {
	pull()
	push()
	handshake()
}
