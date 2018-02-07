package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var workerID int64
var publisherID int64

func main() {
	input := make(chan string)
	done := make(chan bool)
	defer close(done)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input, done)
	go publisher(input, done)
	go publisher(input, done)
	go publisher(input, done)
	time.Sleep(10 * time.Millisecond)
}

// publisher pushes data into a channel
func publisher(out chan string, done chan bool) {
	atomic.AddInt64(&publisherID, 1)
	//publisherID++
	//thisID := publisherID
	thisID := atomic.LoadInt64(&publisherID)
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
		if atomic.LoadInt64(&thisID) > 10 {

			return
		}

	}
}

func workerProcess(in <-chan string) {
	atomic.AddInt64(&workerID, 1)
	//workerID++
	thisID := atomic.LoadInt64(&workerID)
	for n := range in {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := n
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

/*
CHALLENGE #1
Is this fan out?

CHALLENGE #2
Is this fan in?
*/
