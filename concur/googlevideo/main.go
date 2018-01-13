package main

/*
From: https://www.youtube.com/watch?v=f6kdp27TYZs
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func usef(n int) {
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		fmt.Println(i, n)
		left = right
		time.Sleep(time.Second * 2)
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func boring2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			//time.Sleep(time.Second)
			wait := time.Duration(rand.Intn(10000)) * time.Millisecond
			time.Sleep(wait)
			fmt.Println(wait, msg)

		}

	}()
	return c
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		//time.Sleep(time.Second)
		wait := time.Duration(rand.Intn(100000)) * time.Millisecond
		time.Sleep(wait)
		fmt.Println(msg, wait)

	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func fanInOld(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
func main() {
	//	c := make(chan string)
	//	go boring("boring", c)
	//c := boring2("Goodbye")
	//anotherC := boring2("Hello")

	c := fanIn(boring2("Hello"), boring2("Goodbye"))
	c2 := fanIn(boring2("Hello2"), boring2("Goodbye2"))
	go usef(4)
	for i := 0; i < 10; i++ {
		//fmt.Println(<-anotherC)
		go usef(10)
		fmt.Println(<-c)
		fmt.Println(<-c2)
		//fmt.Printf("sleeping: %q\n", <-c)
		select {
		case v1 := <-c:
			fmt.Printf("received %v from c1\n", v1)
		case v2 := <-c2:
			fmt.Printf("received %v from c2\n", v2)
		case <-time.After(1 * time.Second):
			fmt.Println("My old buddy, you're moving much to slow...")
		default:
			fmt.Printf("No one is ready!\n")
		}

	}

}
