package main

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitorForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitorForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitorForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
