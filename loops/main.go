package main

import "fmt"

func main() {
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
