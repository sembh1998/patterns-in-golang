package main

import (
	"fmt"
	"time"
)

func someFunc(entry string) {
	fmt.Println(entry)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(time.Millisecond * 1)

	fmt.Println("hi")
}
