package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func hardIsPrime(num uint64) bool {
	if num == 2 {
		return true
	} else if num%2 == 0 || num == 1 {
		return false
	}
	testTo := uint64(math.Ceil(math.Sqrt(float64(num))))
	for current := uint64(3); current <= testTo; current += 2 {
		if num%current == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	/* 	Input: an integer n > 1.

	Let A be an array of Boolean values, indexed by integers 2 to n,
	initially all set to true.

	for i = 2, 3, 4, ..., not exceeding √n:
	  if A[i] is true:
	    for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n:
	      A[j] := false.

	Output: all i such that A[i] is true.
	*/

	notSmartPrime1 := func(
		done <-chan interface{},
		primeStream <-chan int,
	) <-chan int {
		localStream := make(chan int)
		go func() {
			defer close(localStream)
			localPrime := 0
			for v := range primeStream {
				if hardIsPrime(uint64(v)) {
					localPrime = v
				} else {
					localPrime = 0
				}
				select {
				case <-done:
					return
				default:
					if localPrime != 0 {
						localStream <- localPrime
					}
				}

			}

		}()
		return localStream
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

	fanIn := func(
		done <-chan interface{},
		channels ...<-chan int,
	) <-chan interface{} {
		var wq sync.WaitGroup
		localStream := make(chan interface{})

		multiplex := func(c <-chan int) {
			defer wq.Done()
			for i := range c {
				select {
				case <-done:
					return
				case localStream <- i:
				}
			}
		}
		wq.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		go func() {
			wq.Wait()
			close(localStream)
		}()
		return localStream
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

	done0 := make(chan interface{})
	aLotOfInts := make([]int, 5)
	for i := 5; i < 1000000; i++ {
		aLotOfInts = append(aLotOfInts, i)
	}

	primeStream := generator(done0, aLotOfInts...)
	//	primeStream := generator(done0, 1, 2, 3, 4, 5, 1729, 17, 47)
	numFinders := 10
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = notSmartPrime1(done0, primeStream)
	}

	for p := range take(done0, fanIn(done0, finders...), 100) {
		fmt.Printf("prime: %v\n", p)
	}
	fmt.Printf("This took: %v\n", time.Since(start))

}
