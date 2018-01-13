package main

import (
	"fmt"
	"strconv"
)

func main() {
	g := []byte("Goodbye")
	fmt.Println([]byte("Hello"))
	fmt.Println(g)
	s, _ := strconv.Atoi("6")
	fmt.Println(s)
	for _, b := range g {
		fmt.Sprintln(b)
		fmt.Printf("%v\n", string(int(b)))
	}
	fmt.Println()
}
