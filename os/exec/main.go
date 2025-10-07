package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func run() {
	cmd := exec.Command("ipconfig")

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
}

func wait() {
	cmd := exec.Command("sleep", "3")

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed: %v", err)
		return
	}

	fmt.Println("Command running in the background...")

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Command wait failed： %v", err)
	}
}

func withContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "3")

	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %v\n", err)
	}
}

func outputMsg() {
	cmd := exec.Command("echo", "hello world")

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	fmt.Println(string(output))
}

func main() {
	outputMsg()
}
