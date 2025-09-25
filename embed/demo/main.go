package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed hello.txt
var content string

//go:embed file/2.txt
var contentBytes []byte

//go:embed hello.txt
//go:embed file
var fileFS embed.FS
var data, _ = fileFS.ReadFile("hello.txt")

func main() {
	fmt.Printf("hello.txt content: %s\n", content)

	fmt.Printf("2.txt content: %s\n", contentBytes)

	dir, _ := fs.ReadDir(fileFS, "file")
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Printf("%+v\n", struct {
			Name  string
			IsDir bool
			Info  struct {
				Name string
				Size int64
				Mode fs.FileMode
			}
		}{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Info: struct {
				Name string
				Size int64
				Mode fs.FileMode
			}{Name: info.Name(), Size: info.Size(), Mode: info.Mode()},
		})
	}
}
