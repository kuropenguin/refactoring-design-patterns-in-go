package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance != nil {
		fmt.Println("Single instance already created")
		return singleInstance
	}
	lock.Lock()
	defer lock.Unlock()
	if singleInstance != nil {
		fmt.Println("Single instance already created")
		return singleInstance
	}
	singleInstance = &single{}
	fmt.Println("Create single instance")
	return singleInstance
}
