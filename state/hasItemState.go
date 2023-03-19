package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) despenseItem() error {
	return fmt.Errorf("please select item first")
}
