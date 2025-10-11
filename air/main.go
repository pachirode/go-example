package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Printf("args[1]: %v\n", os.Args[1])

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Air! (PID: %d)\n", os.Getpid())
	})

	go func() {
		fmt.Printf("Server started at http://localhost:8080 (PID %d)\n", os.Getpid())
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	<-signalChan
	fmt.Printf("Server shutting down\n")
	server.Close()
}
