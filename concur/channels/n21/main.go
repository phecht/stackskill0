package main

import (
	"fmt"
	"sync"
)

var wq sync.WaitGroup

func main() {

	myCh := make(chan string)

	wq.Add(2)
	go func() {
		defer wq.Done()
		for i := 0; i < 10; i++ {
			myCh <- "anom 1"
		}
	}()

	go func() {
		defer wq.Done()
		for i := 0; i < 10; i++ {
			myCh <- "anom 2"
		}

		/* 		fmt.Println(<-myCh)
		   		fmt.Println("Clear!") */

	}()

	go func() {
		wq.Wait()
		close(myCh)
		fmt.Println("Closed myCH")

	}()

	for ch := range myCh {
		fmt.Println("ch:", ch)
	}

}
