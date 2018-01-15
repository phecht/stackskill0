package main

import (
	"fmt"
)

/*
why does <-done get called twice?
*/
func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for g := 0; g < 10; g++ {
			c <- g
			fmt.Println("Putting g:", g, " into channel")
		}
		fmt.Println("counted to 10 g")
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("Putting i:", i, " into channel")
		}
		done <- true
	}()

	go func() {
		_ = <-done
		// _ = <-done
		close(c)
		close(done)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
