package main

import (
	"fmt"
	"strconv"
)

//var count int64
//var wg sync.WaitGroup

func main() {
	//wg.Add(2)
	inc1 := incrementor(5)
	var count int
	for counter := range inc1 {
		count++
		fmt.Println(counter)
	}
	fmt.Println("Final Counter:", count)

}

/* 	count := merge(inc1)
}

func merge(cs ...<-chan string) int {
	count := 0
	go func() {
		for _, c := range cs {
			count++
			fmt.Println(c)
		}

	}()
	return count
*/

func incrementor(incQty int) <-chan string {
	out := make(chan string)
	//defer close(out)
	done := make(chan bool)

	for i := 0; i < incQty; i++ {
		// send out a function that will do the adding.
		// putting the results on out
		// putting on the done channel when done.
		go func(i int) {
			for j := 0; j < 20; j++ {
				//atomic.AddInt64(&count, 1)
				outString := "Process: " + strconv.Itoa(i) + " printing:" + strconv.Itoa(j)
				out <- outString
			}
			done <- true

		}(i)

	}

	// this function will pull off the done for each function call above.
	go func() {
		for i := 0; i < incQty; i++ {
			<-done
		}
		close(out)
		close(done)
	}()
	return out
	//wg.Done()
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
