package main

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitorForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
