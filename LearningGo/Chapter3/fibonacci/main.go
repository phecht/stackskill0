package main

import "fmt"

func main() {
	counter := 40
	fib(counter)
	// evidently my original func didn't use slices.
	fmt.Println(fibGo(10))
}
func fib(lastone int) {
	prevNum := 0
	thisNum := 0
	for numb := 1; numb <= lastone; numb++ {
		switch numb {
		case 1:
			fmt.Println(numb)
			prevNum = 1
			thisNum = 1
		case 2:
			fmt.Println(numb)
			prevNum = 1
			thisNum = 2
		default:
			nextNum := thisNum + prevNum
			fmt.Println(nextNum)
			prevNum = thisNum
			thisNum = nextNum
		}
	}
}

func fibGo(fibCounter int) []int {
	fibList := make([]int, fibCounter)
	if fibCounter == 0 {
		return fibList
	}
	if fibCounter == 1 {
		fibList[0] = 1
		return fibList
	}
	if fibCounter == 2 {
		fibList[0] = 1
		fibList[1] = 1
		return fibList
	}
	for thisFib := 2; thisFib < fibCounter; thisFib++ {
		fibList[0] = 1
		fibList[1] = 1
		fmt.Println("Here:", thisFib, "fibs:", fibList[thisFib-1], fibList[thisFib-2])
		fibList[thisFib] = fibList[thisFib-1] + fibList[thisFib-2]
	}
	return fibList
}
