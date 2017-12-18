package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	index int
	data  [20]int
}

func main() {

	var mainStack stack
	fmt.Println("mainStack.index:", mainStack.index)
	mainStack.index = 0
	mainStack.push(14)
	mainStack.push(9)
	var sortSlice []int
	sortSlice = make([]int, 10)
	p2 := plusTwo()
	p3 := plusX(9)
	fmt.Printf("%d\n", p2(9))
	fmt.Printf("%d\n", p3(10))
	fmt.Println("Slice sort")
	fmt.Scanln()
	setSortSlice(sortSlice)
	fmt.Println("sortSlice", sortSlice)
	fmt.Println("max", maxSortSlice(sortSlice))

	// Use sortSlice and p3 for a map function
	fmt.Printf("%v\n", (mapx(p3, sortSlice)))
	fmt.Printf("mainStack String: %v\n", mainStack.String())

	fmt.Printf("mainStack: %v\n", mainStack)
	what := mainStack.pop()
	fmt.Printf("What: %v\n", what)
	fmt.Printf("mainStack: %v\n", mainStack)
	fmt.Printf("mainStack String: %v\n", mainStack.String())
	what = mainStack.pop()

	fmt.Printf("What: %v\n", what)
	fmt.Printf("mainStack: %v\n", mainStack)
	what = mainStack.pop()
	fmt.Printf("What: %v\n", what)
	fmt.Printf("mainStack: %v\n", mainStack)

}

func (myStack stack) String() string {
	var str string
	for i := 0; i <= myStack.index; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(myStack.data[i]) + "]"
	}
	return str
}

func (myStack *stack) push(slot int) {
	myStack.data[myStack.index] = slot
	myStack.index++
}

func (myStack *stack) pop() int {
	if myStack.index < 1 {
		return 0
	}
	myStack.index--
	slot := myStack.data[myStack.index]
	return slot
}

func mapx(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func plusTwo() func(int) int {
	return func(x int) int { return x + 2 }
}

func plusX(x int) func(int) int {
	return func(y int) int { return x * y }
}
func setSortSlice(ss []int) {
	ss[0] = 10
	ss[1] = 45
	ss[2] = 765
	ss[3] = 1729
	ss[4] = 23
	ss[5] = 56
	ss[6] = 90
	ss[7] = 121
	ss[8] = 424
	ss[9] = 900
}

func maxSortSlice(ss []int) (maxValue int) {
	maxValue = 0
	for _, thisVal := range ss {
		if thisVal > maxValue {
			maxValue = thisVal
		}
	}
	return
}
