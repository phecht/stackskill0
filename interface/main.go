package main

import (
	"fmt"
	"math"
	"math/big"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (cir Circle) area() float64 {
	return math.Pi * math.Pow(cir.radius, 2)
}

type Shaper interface {
	area() float64
}

func info(s Shaper) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	sq := Square{10}
	ci := Circle{5}
	fmt.Println("-------------------")
	fmt.Println(sq.area())
	fmt.Println("-------------------")
	info(sq)
	info(ci)
	n := big.NewInt(0)
	fmt.Println(n.SetBit(n, 100, 1))
	z := new(big.Int).Exp(big.NewInt(2), big.NewInt(2), nil)
	fmt.Printf("%T\n", *z)
	fmt.Println(z)
}
