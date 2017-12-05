package main

import "fmt"

func zzero(z *int) {
	fmt.Println(z)
	*z = 0
}
func zero(x int) {
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
	zzero(&x)
	fmt.Println(x)

}
