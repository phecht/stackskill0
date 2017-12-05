package main

import "fmt"

func main() {
	counter := 1

Start:
	fmt.Println(counter)
	counter++
	if counter <= 10 {
		goto Start
	}

}
