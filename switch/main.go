package main

import "fmt"

func main() {
	fmt.Println("Switch")
	switchOnType(9)
	switchOnType("foo")
	switchNoParam()
}

func switchNoParam() {
	myInt := 10
	myString := "Mu"
	switch {
	case myInt == 10:
		fmt.Println(myInt)
		fallthrough
	case myString == "Mu":
		fmt.Println("Mu is the string", len(myString))
		fallthrough
	case len(myString) == 2:
		fmt.Println(myString, " ", len(myString))
	case len(myString) == 2:
		fmt.Println(myString, "-=", len(myString))
	}
}
func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("String")
	}
}
