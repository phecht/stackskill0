package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f string
	f = ""
	fmt.Println("Enter number:")
	fmt.Scanln(&f)
	i, err := strconv.ParseInt(f, 10, 32)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("Factorial for", i)
	fmt.Println("Old school", factorial(int(i)))
	c := chanFact(int(i))
	for x := range c {
		fmt.Println(x)
	}
}

func chanFact(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
