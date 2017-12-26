package main

import (
	"fmt"
	"math"
)

/* based on https://en.wikipedia.org/wiki/Newton%27s_method
and https://tour.golang.org/flowcontrol/8

the round and toFIxed are modified from https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision-in-golang

*/
func sqrt(value, guess float64) float64 {
	guess -= (guess*guess - value) / (2 * guess)
	return guess
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

/* evidently only works to precision of 11
 */
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	guess := 1.0
	result := 1.0
	square := 40073467.0
	myprecision := 11
	for i := 1; i < 200; i++ {
		result = guess

		guess = sqrt(square, guess)
		fmt.Println("Guess", toFixed(guess, myprecision))
		if toFixed(guess, myprecision) == toFixed(result, myprecision) {
			fmt.Println(toFixed(guess, myprecision))
			return

		}
	}
}
