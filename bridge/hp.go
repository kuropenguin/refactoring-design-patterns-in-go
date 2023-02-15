package main

import "fmt"

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("printing by a HP printer")
}
