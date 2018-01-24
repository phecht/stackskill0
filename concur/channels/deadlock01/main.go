package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// Wrapping sending 1 to c in an anonynous go function removes the deadlock.
	go func() {
		c <- 1
		fmt.Println("func")

	}()

	//
	fmt.Println("Here")
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?
