package main

import "fmt"

func main() {
	list1 := []float64{3.56, 5.88, 88.900, 2.45678}
	fmt.Println("Average is:", avg1(list1))

}
func avg1(list []float64) (avg float64) {
	var total float64
	total = 0.0
	switch len(list) {
	case 0:
		total = 0
	default:
		for _, val1 := range list {
			total += val1
		}
		avg = total / float64(len(list))
	}
	return
}
