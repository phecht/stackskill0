package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
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
			fmt.Println("Summing ", sum, " and ", n, " to sum.")
			sum += n
		}
		fmt.Println("Putting ", sum, " on channel out.")
		out <- sum
		close(out)
	}()
	return out
}
