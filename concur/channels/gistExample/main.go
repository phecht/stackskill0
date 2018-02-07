package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var workerID int64
var publisherID int64

// Publisher push data into a channel
func Publisher(out chan string) {
	//publisherId += 1
	atomic.AddInt64(&publisherID, 1)
	thisID := atomic.LoadInt64(&publisherID)
	dataID := 0
	for {
		dataID++
		fmt.Printf("Publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

// WorkerProcess takes the in chan and put it on the input channel
func WorkerProcess(in <-chan string) {
	//workerID += 1
	atomic.AddInt64(&workerID, 1)
	thisID := atomic.LoadInt64(&workerID)
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

func main() {
	input := make(chan string)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	time.Sleep(1 * time.Millisecond)
}
