package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		fmt.Println("making cake", cakeName)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbryCS chan string, chocCS chan string) {
	closedStraw, closedChoc := false, false

	for {
		//if both channels are closed then we can stop
		if closedStraw && closedChoc {
			return
		}
		// 	fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, okStraw := <-strbryCS:
			if !okStraw && !closedStraw {
				closedStraw = true
				fmt.Println(" ... Strawberry channel closed!")
			} else if !closedStraw {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, okChoc := <-chocCS:
			if !okChoc && !closedChoc {
				closedChoc = true
				fmt.Println(" ... Chocolate channel closed!")
			} else if !closedChoc {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}

func main() {
	strawCakeCS := make(chan string)
	chocCakeCS := make(chan string)

	cakeCount := 7
	//two cake makers
	go makeCakeAndSend(chocCakeCS, "Chocolate", cakeCount)   //make3 cakeCount chocolate cakes and send
	go makeCakeAndSend(strawCakeCS, "Strawberry", cakeCount) //make cakeCount strawberry cakes and send

	//one cake receiver and packer
	go receiveCakeAndPack(strawCakeCS, chocCakeCS) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(1 * time.Second)
}
