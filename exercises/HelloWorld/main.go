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
	fmt.Println("Press a key for listNum..")
	fmt.Scanln()
	listNum()
	fmt.Println("Press a key for fizzBuzz..")
	fmt.Scanln()
	fizzBuzz()
	fmt.Println("Press a key for fizzCount..")
	fmt.Scanln()
	fizzCount()
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
	fmt.Printf("Small %v, Large %v, remainder %v\n", smallNum, largeNum, theRemander)
}

func listNum() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func fizzBuzz() {

	for fb := 1; fb <= 100; fb++ {
		normal := true
		if fb%3 == 0 {
			fmt.Printf("FIZZ")
			normal = false
		}
		if fb%5 == 0 {
			fmt.Printf("BUZZ")
			normal = false
		}
		if normal {
			fmt.Printf("%v", fb)
		}
		fmt.Println()
	}
}

func fizzCount() {

	var total35 int
	for fb := 1; fb < 1000; fb++ {

		if fb%3 == 0 || fb%5 == 0 {
			total35 += fb
		}
	}
	fmt.Printf("Total is %d\n", total35)
}
