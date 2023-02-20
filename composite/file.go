package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyboard string) {
	fmt.Printf("Search for keyword %s in file %s\n", keyboard, f.name)
}

func (f *File) getName() string {
	return f.name
}
