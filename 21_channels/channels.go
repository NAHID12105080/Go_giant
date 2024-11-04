package main

import (
	"fmt"
	"math/rand"
)

func processNum(chageNumber chan int){
	for num:= range chageNumber{
		fmt.Println("processing number", num)
		// time.Sleep(time.Second)
	}
}


func main() {
	// creating a channel
	// messages := make(chan string)

	// // sending a value into a channel
	// go func() { messages <- "ping" }()

	// // receiving a value from a channel
	// msg := <-messages
	// fmt.Println(msg)

	// time.Sleep(time.Second*2)
	//-----------basic channel example--------
	
	chageNumber:=make(chan int)
	//sending data
	go processNum(chageNumber)
	for{
		chageNumber <-rand.Intn(100)	
	}
}