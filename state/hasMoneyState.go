package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) despenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount -= 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
