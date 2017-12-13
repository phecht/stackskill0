package main

import "fmt"

func main() {
	theList := []int{3, 77, 23, 10, 60, 55, 1, 5, 99}
	fmt.Println("unsorted:", theList)
	bubbleSort(theList)
	fmt.Println("sorted:", theList)

}

func bubbleSort(list []int) {
	listlen := len(list)
	fmt.Println("List length:", listlen)
	switch listlen {
	case 0:
		return
	default:
		for start := 0; start < listlen-1; start++ {
			for bubble := start + 1; bubble < listlen; bubble++ {
				if list[bubble] < list[start] {
					fmt.Print("Swapping position ", bubble, " with position ", start, ", ")
					fmt.Println("Swapping value ", list[bubble], " with position ", list[start])
					swap := list[bubble]
					list[bubble] = list[start]
					list[start] = swap

				}
			}
		}

	}
}
