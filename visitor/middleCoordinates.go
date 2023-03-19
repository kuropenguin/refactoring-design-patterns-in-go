package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitorForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitorForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitorForRectangle(r *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
