package main

import "fmt"

func main() {
	for iterator := 32; iterator < 20000; iterator++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", iterator, iterator, iterator, iterator)
	}
}
