package main

import (
	"fmt"
	"os"
)

func OpenFile(path string) (*os.File, func(), error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		fmt.Println("cleanup...")
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}

	return f, cleanup, nil
}

type Content string

func Clean() (Content, func(), error) {

	cleanup := func() {
		fmt.Println("cleanup...")
	}

	return "clean", cleanup, nil
}

type App struct {
	File    *os.File
	Cleanup Content
}
