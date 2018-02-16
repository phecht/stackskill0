package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"text/tabwriter"
	"time"
)

func myDebug(flag bool, dstring ...interface{}) {
	if flag {
		retString := append([]interface{}{}, dstring...)
		fmt.Println(retString)
	}
}

func main() {

	const debug = true
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			myDebug(debug, "Producing:"+strconv.Itoa(i))
			l.Lock()
			l.Unlock()
			time.Sleep(1)

		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		defer l.Unlock()
		l.Lock()
		myDebug(debug, "Leaving observer...")

	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		// We are going to add count+1 waits
		// We are going to run count observer gos
		// We are going to run the producer, so count+1
		wg.Add(count + 1)
		myDebug(debug, "Count:"+strconv.Itoa(count+1))

		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			myDebug(debug, "Running observer with "+strconv.Itoa(i))
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}
