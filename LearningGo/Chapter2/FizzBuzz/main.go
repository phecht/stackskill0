package main

import "fmt"

func main() {

	for fb := 1; fb <= 100; fb++ {
		normal := true
		if fb%3 == 0 {
			fmt.Printf("FIZZ")
			normal = false
		}
		if fb%5 == 0 {
			fmt.Printf("BUZZ")
			normal = false
		}
		if normal {
			fmt.Printf("%v", fb)
		}
		fmt.Println()
	}
}
