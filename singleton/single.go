package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creatin instance")
			singleInstance = &single{}
		} else {
			fmt.Println("instance already created")
		}
	} else {
		fmt.Println("instace already exists")
	}
	return singleInstance
}

func main() {
	for range 10 {
		go getInstance()
	}

	//to doesn't let the main gourutine die and the other don't excecute
	fmt.Scanln()
}
