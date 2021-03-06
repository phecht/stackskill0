package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

func myWalker(t *tree.Tree) {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	if t.Left != nil {
		myWalker(t.Left)
	}
	if t.Right != nil {
		myWalker(t.Right)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func myString(t *tree.Tree) string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += myString(t.Left) + " "
		fmt.Println("Adding left " + s)
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + myString(t.Right)
		fmt.Println("Adding right " + s)
	}
	return "(" + s + ")"
}
func main() {

	t23 := tree.New(1)
	t2 := tree.New(2)
	t3 := tree.New(3)
	t23.Left = t2
	t23.Right = t3

	tt23 := tree.New(1)
	tt2 := tree.New(2)
	tt3 := tree.New(3)
	tt23.Left = tt2
	tt23.Right = tt3
	fmt.Println(tt23.String())
	fmt.Println(t23.String())

	fmt.Println("t23")
	myWalker(t23)
	fmt.Println("tt23 myWalker")
	myWalker(tt23)
	fmt.Println("t23")
	myString(t23)

	fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))
	//this test fails as each tree.New() is a random generated tree of 10 nodes
	fmt.Println("t23 and tt23 same: ", Same(t23, tt23))
}
