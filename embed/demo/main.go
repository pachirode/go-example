package main

import "embed"

//go:embed hello.txt
var content string

//go:embed hello.txt
var contentBytes []byte

//go:embed hello.txt
var fileFS embed.FS
var data, _ = fileFS.ReadFile("hello.txt")
