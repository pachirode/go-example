package foo

import (
	_ "unsafe"

	_ "github.com/pachirode/go-example/build/linkname/bar"
)

// Pull

//go:linkname Test github.com/pachirode/go-example/build/linkname/bar.test
func Test()

// Push

func TestPush()

// Handshake

//go:linkname TestHandshake github.com/pachirode/go-example/build/linkname/bar.testHandshake
func TestHandshake()
