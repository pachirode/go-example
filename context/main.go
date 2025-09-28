package main

import "fmt"

func main() {
	var closedchan = make(chan struct{})
	fmt.Printf("%v", closedchan)

	close(closedchan)
	fmt.Printf("%v", closedchan)
}
