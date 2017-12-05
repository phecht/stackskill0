package main

import "fmt"

func main() {
	var arr [10]int

	for mycounter := 0; mycounter < 10; mycounter++ {
		arr[mycounter] = mycounter
	}

	// evidently a dumb way to do it!
	/*
		for mycounter := 0; mycounter < 10; mycounter++ {
			fmt.Println(arr[mycounter])
		}
	*/
	fmt.Printf("%v", arr)
}
