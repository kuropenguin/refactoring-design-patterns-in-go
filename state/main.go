package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.despenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.despenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
