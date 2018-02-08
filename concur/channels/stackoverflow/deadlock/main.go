package main

import (
	"fmt"
)

func main() {
	_gen := generator()
	_inChannel := _inListener(_gen)
	for val := range _inChannel {
		fmt.Print(val, " -- \n")
	}
}

func generator() chan int { // NEED TO CALCULATE FACTORIAL FOR 100 NUMBERS
	ch := make(chan int) // CREATE CHANNEL TO INPUT NUMBERS
	done := make(chan bool)
	go func() {
		defer close(ch)
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		//close(ch) // CLOSE CHANNEL WHEN ALL NUMBERS HAVE BEEN WRITTEM
		done <- true
	}()

	// this function will pull off the done for each function call above.
	go func() {
		for i := 1; i < 100; i++ {
			<-done
		}
	}()

	return ch
}

func _inListener(ch chan int) chan int {
	rec := make(chan int) // CHANNEL RECEIVED FROM GENERATOR
	go func() {
		for num := range ch { // RECEIVE THE INPUT NUMBERS FROM GENERATOR
			result := factorial(num) // RESULT IS A NEW CHANNEL CREATED
			rec <- result            // MERGE INTO A SINGLE CHANNEL; rec
		}
		close(rec)
	}()

	return rec // RETURN THE DEDICATED CHANNEL TO RECEIVE ALL OUTPUTS
}

func factorial(n int) int {
	// OF FACTORIAL
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total // RETURN THE CHANNEL HAVING THE FACTORIAL CALCULATED
}
