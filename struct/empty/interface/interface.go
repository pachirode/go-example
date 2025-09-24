package main

type Writer interface {
	Write(p []byte) (n int, err error)
}

var Discard Writer = discard{}

type discard struct{}

func (discard) Write(p []byte) (int, error) {
	return len(p), nil
}
