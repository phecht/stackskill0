package main

import "fmt"

func main() {

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
	fmt.Println("Ctrl+C")
	fmt.Scanln()

	fmt.Println([]byte("Hello 世界 "))
	fmt.Scanln("Hello")
	fmt.Scanln()

	for chars := 50; chars < 140; chars++ {
		fmt.Println(chars, " - ", string(chars), " -", []byte(string(chars)))
	}
	fmt.Println("Ctrl+C")
	fmt.Scanln()

	for outer := 1; outer <= 10; outer++ {
		for inner := 1; inner <= 10; inner++ {
			fmt.Println(outer, " - ", inner)
		}
	}
	max := 0

	for max < 10 {
		fmt.Println(max)
		max++
	}

	for {
		fmt.Println(max)
		if max >= 2000 {
			break
		}
		max++
	}

	fmt.Scanln()

	max = 0
	for {
		max++
		if max%2 == 0 {
			continue
		}
		fmt.Println(max)

		if max >= 50 {
			break
		}
	}

	fmt.Scanln()

	for {
		fmt.Println(max)
		max++
	}

}
