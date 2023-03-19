package main

type Visitor interface {
	visitorForSquare(*Square)
	visitorForCircle(*Circle)
	visitorForRectangle(*Rectangle)
}
