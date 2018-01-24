package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	/*
	   	go func() {
	   		for i := 0; i < 10; i++ {
	   			fmt.Println(<-c)
	   		}
	   	}()
	    	foo := ""
	   	fmt.Scanln(&foo)
	*/
	for i := range c {
		fmt.Println(i)

	}
}

// Why does this only print zero?
// The first anon function can only put one vaue in the channel.

// And what can you do to get it to print all 0 - 9 numbers?
// Make the Println into a function that takes a number of for each one put on.
// Range removes the need to have the scanln. Since range goes over each one.
