package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int64)

	fv := []int{}
	fv = make([]int, 100)
	for i := 0; i < 22; i++ {
		fv[i] = i
	}

	startF(m, fv...)
	fmt.Printf("Actual Map of Factorials:\n%v\n\n", m)

	fmt.Println("Output by List of Input Values:")
	for _, r := range fv {
		fmt.Printf("%v:%v\n", r, m[r])
	}
}

func startF(m map[int]int64, fs ...int) {
	out := make(chan int64)
	defer close(out)

	for _, tn := range fs {
		go pFactor(tn, out)
		m[tn] = <-out
	}
}

func pFactor(tn int, out chan int64) {
	result := int64(1)
	for i := 1; i <= tn; i++ {
		result = result * int64(i)
	}
	out <- int64(result)
}
