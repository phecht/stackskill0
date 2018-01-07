package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int      { return len(p) }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) More(i, j int) bool { return p[i] > p[j] }

//func (p people) Swap

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println("Sort:")
	fmt.Println("---- numbers -----")
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	fmt.Println("---- numbers reversed -----")
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
	fmt.Println("---- strings -----")
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println("---- strings reversed -----")
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	//sort.Strings(sort.Reverse(sort.Slice(studyGroup)))
	fmt.Println("---- People -----")
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
