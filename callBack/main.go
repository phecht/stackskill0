package main

import "fmt"

func filter(numbers []int, acallback func(int) bool) []int {
	theSlice := []int{}
	for _, num := range numbers {
		if acallback(num) {
			theSlice = append(theSlice, num)
		}
	}
	return theSlice
}

func main() {
	x := 1
	mainSlice := filter([]int{3, 6, 23, 14, 2, 89, 34}, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(x)
	fmt.Println(mainSlice)
}
