package main

import "fmt"

func base() {
	var closedchan = make(chan struct{})
	fmt.Printf("%v", closedchan)

	close(closedchan)
	fmt.Printf("%v", closedchan)
}

func main() {
	//keyConflict()
	keyConflictFix()
}
