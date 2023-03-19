package main

type Share interface {
	getType() string
	accept(Visitor)
}
