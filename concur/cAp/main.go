package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wq sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wq.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
		time.Sleep(time.Duration(9 * time.Millisecond))
	}
	wq.Done()
}

func main() {
	wq.Add(2)
	go foo()
	go bar()
	wq.Wait()
}
