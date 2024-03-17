package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 1)
		myChannel <- "data"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
