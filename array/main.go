package main

import "fmt"

func arraySkill3() {
	var nums [256]byte

	for ind := 0; ind < len(nums); ind++ {
		nums[ind] = byte(ind)
	}
	for ordinal, value := range nums {
		fmt.Printf("%v - %T - %b -- %v\n", value, value, value, ordinal)
		if ordinal > 50 {
			break
		}
	}
}
func arraySkill2() {
	var nums [256]int

	for ind := 0; ind < len(nums); ind++ {
		nums[ind] = ind
	}
	for ordinal, value := range nums {
		fmt.Printf("%v - %T - %b -- %v\n", value, value, value, ordinal)
		if ordinal > 50 {
			break
		}
	}
}
func arraySkill() {
	var charX [58]string
	fmt.Println(charX)
	fmt.Println(len(charX))
	fmt.Println(charX[34])
	for ind := 65; ind <= 122; ind++ {
		charX[ind-65] = string(ind)
	}
	fmt.Println(charX)
	fmt.Println(len(charX))
	fmt.Println(charX[34])

}
func main() {

	arraySkill()
	fmt.Println("Continue to 2")
	fmt.Scanln()
	arraySkill2()

	fmt.Println("Continue to 3")
	fmt.Scanln()
	arraySkill3()
}
