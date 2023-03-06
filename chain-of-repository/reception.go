package main

import (
	"fmt"
)

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registeration done")
		r.next.execute(p)
		return
	}

	fmt.Println("Reception registring patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
