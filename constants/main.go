package main

import "fmt"

// https://blog.golang.org/constants
// Evidently const can be declared and never used.
const z = "Boring"

const (
	a = iota // 0?

	c = iota // next iota value
	b = iota // depends on placement for value.
)

const (
	_ = iota
	e = iota // 0?

	f = iota // next iota value
	g        // depends on placement for value.
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	const p = 47

	fmt.Println("f -", f)
	fmt.Println("p -", p)
	fmt.Println("a -", a)
	fmt.Println("b -", b)
	fmt.Println("c -", c)
	fmt.Println("e -", e)
	fmt.Println("f -", f)
	fmt.Println("g -", g)
	fmt.Printf("KB:%b\t - %d\n", kb, kb)
	fmt.Printf("MB:%b\t - %d\n", mb, mb)
	fmt.Printf("GB:%b\t - %d\n", gb, gb)
	fmt.Println("mb -", mb)
	fmt.Println("From go blog on constants:")
	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)

	const Huge = 1e1000
	fmt.Println(Huge / 1e998)

}
