package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[Debug] - ", log.Lshortfile)
	logger.Println("logger")

	logger.SetPrefix("[Info] - ")
	logger.Println("logger")
}
