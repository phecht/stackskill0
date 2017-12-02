package main

import "fmt"

func main() {
	a := 0
	b := "golang"
	c := 4.17
	d := false

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
}
