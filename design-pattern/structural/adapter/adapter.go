package adapter

import "fmt"

type IProcess interface {
	Process()
}

type Converter interface {
	Convert()
}

type adaptee struct {}
func (adaptee) Convert() {
	fmt.Println("Converting")
}

type adapter struct {
	c Converter
}

func (a adapter) Process() {
	fmt.Println("Processing")
	a.c.Convert()
}
