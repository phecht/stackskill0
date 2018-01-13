package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wq sync.WaitGroup
var mutex sync.Mutex

var badCounter int
var goodCounter int
var atomicCounter int64

func atomicInc(msg string) {
	for ai := 0; ai < 20; ai++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		atomic.AddInt64(&atomicCounter, int64(1))
		fmt.Println(msg, "Counter:", atomicCounter, "ai:", ai)
	}
	wq.Done()

}

func goodInc(msg string, p bool) {
	for gi := 0; gi < 20; gi++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		x := goodCounter
		x++
		goodCounter = x
		if p {
			fmt.Println(msg, "Counter:", goodCounter, "gi:", gi)
		}
		mutex.Unlock()
	}
	wq.Done()

}

func badInc(msg string, p bool) {
	for i := 0; i < 20; i++ {
		x := badCounter
		x++
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		badCounter = x
		if p {
			fmt.Println(msg, "Counter:", badCounter, "i:", i)

		}
	}
	wq.Done()

}

func main() {
	fmt.Println("race conditions:")
	wq.Add(4)
	go badInc("First", false)
	go badInc("second", false)
	go goodInc("Good first", true)
	go goodInc("Good second", true)
	go atomicInc("Atomic first")
	go atomicInc("Atomic second")

	wq.Wait()
	fmt.Println("Final:b", badCounter)
	fmt.Println("Final g:", goodCounter)
	fmt.Println("Final a:", atomicCounter)

}
