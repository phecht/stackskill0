package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	fmt.Println("linkedlist:")
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(4)

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
