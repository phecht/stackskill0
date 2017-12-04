package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J
			}
			fmt.Println(i)
		}
	}
}
