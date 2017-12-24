package main

import "fmt"

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
const calcNum = 20

func main() {
	fmt.Println("problem5")
	if divideAll(12) {
		fmt.Println("12 is good!")
	}
	fmt.Println(findMax())
	fmt.Println(findLowest(findMax()))
}

func findMax() (max int) {
	max = 1
	for i := 1; i <= calcNum; i++ {
		max = (max * i)
	}
	return max
}
func findLowest(max int) (lowest int) {
	for i := calcNum; i < max; i++ {
		// fmt.Printf("Trying %d\n", i)
		if divideAll(i) {
			return i
		}
	}
	return 0
}
func divideAll(testNum int) (passed bool) {

	for i := 1; i <= calcNum; i++ {
		if testNum%i != 0 {
			return false
		}
	}
	return true
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

I could have sped this up by at least not trying number that are known to not work.

Calculate the largest possible number if they are all prime factors.

Go through each number and divide by all numbers upto calcNum.
*/
