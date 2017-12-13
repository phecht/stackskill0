package main

import "fmt"

func main() {

	for mycounter := 0; mycounter < 10; mycounter++ {
		printnum(mycounter)
	}

}
func printnum(i int) {
	fmt.Println(i)
}
