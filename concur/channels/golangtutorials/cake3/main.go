package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		fmt.Println("Making cake:", cakeName)
		cs <- cakeName //send a strawberry cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs, 6)
	go receiveCakeAndPack(cs)

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(3 * time.Second)
}
