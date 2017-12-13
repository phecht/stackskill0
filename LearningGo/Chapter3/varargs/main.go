package main

import "fmt"

func main() {
	varargs(1, 45, 3, 39)
	varargs(45, 88, 90, 34, 12)

}

func varargs(numberList ...int) {
	for _, num := range numberList {
		fmt.Printf("%d\n", num)
	}
}
