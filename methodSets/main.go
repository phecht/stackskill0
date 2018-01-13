package main

import (
	"fmt"
	"math"
)

type circlePointer struct {
	radius float64
}
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (cP *circlePointer) area() float64 {
	return math.Pi * math.Pow(cP.radius, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Printf("area %v\n", s.area())
}

func main() {
	c := circle{5}
	info(&c)
	info(c)
	cP := circlePointer{4}
	info(&cP)
	// info(c) won't work since we need to pass a pointer.
}
