package main

import (
	"fmt"
	"sync"
)

var wq sync.WaitGroup

func main() {

	myCh := make(chan string)

	wq.Add(3)
	go func() {
		for i := 0; i < 10; i++ {
			myCh <- "anom 1"
		}
		wq.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			myCh <- "anom 2"
		}

		/* 		fmt.Println(<-myCh)
		   		fmt.Println("Clear!") */
		wq.Done()
	}()

	go func() {
		wq.Wait()
		close(myCh)
		wq.Done()
	}()

	for ch := range myCh {
		fmt.Println("ch:", ch)
	}

}
