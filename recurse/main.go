package main

import "fmt"

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	//mt.Println("x:", x)
	return x * factorial(x-1)
}

func changeMe(z map[string]int) {
	z["Peter"] = 18
}

func main() {
	var inputNum int64

	m := make(map[string]int)
	fmt.Println("Enter a number to factorize (up to 20:")
	fmt.Scan(&inputNum)
	fmt.Println(factorial(inputNum))
	changeMe(m)
	fmt.Println(m["Peter"])
}
