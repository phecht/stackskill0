package main

import (
	"fmt"
	"math/rand"
)

func main() {

	repeatFN := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		localStream := make(chan interface{})
		go func() {
			defer close(localStream)
			for {
				select {
				case <-done:
					return
				case localStream <- fn():
				}
			}

		}()
		return localStream

	}

	repeat := func(
		done <-chan interface{},
		values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream

	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}

		}()
		return takeStream

	}

	generator := func(
		done <-chan interface{},
		integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

	done2 := make(chan interface{})
	defer close(done2)
	for num := range take(done2, repeat(done2, 10), 10) {
		fmt.Printf("%v\n", num)
	}

	done3 := make(chan interface{})
	defer close(done3)

	// This seems to return the same random numbers.  Probably need a seed.
	rand := func() interface{} { return rand.Int() }
	for num := range take(done3, repeatFN(done3, rand), 10) {
		fmt.Printf("%v\n", num)
	}

}
