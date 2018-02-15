package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	inStream := make(chan string)

	terminated := doWork(done, inStream)

	go func() {
		// rare comment
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}
