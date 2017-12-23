package main

import "fmt"

/* func main() {
	fmt.Println("Hello World")
} */

func main() {
	var name string
	var err error
	fmt.Printf("Type your name and hit return:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Some error occured")
	}
	fmt.Printf("Hello %v\n", name)
	remainer()
}

func remainer() {
	var smallNum, largeNum, theRemander int
	var err error

	fmt.Printf("Smaller number:")
	_, err = fmt.Scanf("%d", &smallNum)
	if err != nil {
		fmt.Println("Some error occured")
		return
	}
	fmt.Printf("Larger number:")
	_, err = fmt.Scanf("%d", &largeNum)
	if err != nil {
		fmt.Println("Some error occured")
		return
	}
	if smallNum >= largeNum {
		fmt.Println("Run again and please have smaller number, actually smaller than the larger number!")
	}
	theRemander = largeNum - (smallNum % largeNum)
	fmt.Printf("Small %v, Large %v, remainder %v", smallNum, largeNum, theRemander)
}
