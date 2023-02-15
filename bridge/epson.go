package main

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("printing by a epson printer")
}
