package main

import (
	"fmt"
	"strconv"
	"time"
)

var cakeMake int

func makeCakeAndSend(cs chan string) {
	for i := 1; i <= cakeMake; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		fmt.Println("Making a cake and sending ...", cakeName)
		cs <- cakeName //send a strawberry cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for i := 1; i <= cakeMake; i++ {
		s := <-cs //get whatever cake is on the channel
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cakeMake = 6
	cs := make(chan string)
	go makeCakeAndSend(cs)
	go receiveCakeAndPack(cs)

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(4 * time.Second)

}
