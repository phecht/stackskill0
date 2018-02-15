package main

import (
	"fmt"
	"sync"
)

func main() {
	var wq sync.WaitGroup
	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wq.Add(1)
		go func(s string) {
			defer wq.Done()
			//fmt.Println("WTF")
			fmt.Println(s)
		}(salutation)
		fmt.Println(salutation)
	}
	wq.Wait()
}
