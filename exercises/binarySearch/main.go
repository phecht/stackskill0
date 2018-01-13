package main

import (
	"fmt"
)

func main() {
	list := []int{1, 5, 9, 12, 30, 50, 78, 345, 5678}
	fmt.Println(list)
	fmt.Println(BinarySeach(list, 501))
}
func BinarySeach(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}

		}
	}
	return false
}
