package main

import "fmt"

func main() {
	total := 0.0

	var a [5]float64
	var myavg float64
	myavg = 0
	a[0] = 4.5
	a[1] = 6.78
	a[2] = 3.56
	a[3] = 2.138
	a[4] = 4.09

	myslice := a[0:4]

	switch len(myslice) {
	case 0:
		myavg = 0
	default:
		for _, v := range myslice {
			total += v
		}
		myavg = total / float64(len(myslice))
	}

	fmt.Println("Average:", myavg)
}
