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
	fmt.Println("Press a key for Half..")
	fmt.Scanln()
	fmt.Println(half(3))
	h := func(i int) (r int, e bool) {
		return i / 2, i%2 == 0
	}
	fmt.Println(h(6))

	fmt.Println((true && false) || (false && true) || !(false && false))
	fmt.Println("Press a key for largestInt..")
	fmt.Scanln()
	nums := []int{3, 10, 50, 10, 1762, 1729, 4}
	largestInt(nums...)
	fmt.Println("Press a key for foo..")
	fmt.Scanln()
	foo(1, 2)
	foo(1, 2, 3)
	foo(nums...)

}

func foo(numberList ...int) {
	for _, num := range numberList {
		fmt.Printf("%d\n", num)

	}
}

func largestInt(numberList ...int) {
	var lI int
	lI = 0
	for _, num := range numberList {
		if num > lI {
			lI = num
		}

	}
	fmt.Printf("%d\n", lI)
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
	theRemander = (largeNum % smallNum)
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

func half(sum int) (x int, even bool) {
	x = sum / 2
	even = (sum%2 == 0)
	return x, even
}
