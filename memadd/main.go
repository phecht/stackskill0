package main

import "fmt"

func main() {
	a := 4
	var meters, yards float64
	var b *float64
	const m2y float64 = 1.09361
	fmt.Println("a - ", a)
	fmt.Println("a address - ", &a)
	fmt.Printf("%d  \n", &a)
	fmt.Println("Enter meters:")
	fmt.Scan(&meters)
	yards = meters * m2y
	fmt.Println(meters, " meters is", yards, " yards!")
	b = &meters
	fmt.Println(&meters, " is ", b)
	fmt.Println(*b)
	*b = 300
	fmt.Printf("Meters is now %v\n", meters)

}
