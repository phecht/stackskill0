package main

import (
	"fmt"
)

func main() {
	chan1 := incrementor("Foo:")
	chan2 := incrementor("Bar:")
	c3 := puller(chan1)
	c4 := puller(chan2)
	fmt.Println("Final X Counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 40; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
