package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			c <- i
			fmt.Println("set channel")
		}
		close(c)
	}()

	/* 	go func() {
		for n := range c {
			x := <-c
			fmt.Println("func", x, n)

		}

	}() */
	for n := range c {
		x := <-c
		fmt.Println(n, x)

	}

	// time.Sleep(time.Second * 2)
}
