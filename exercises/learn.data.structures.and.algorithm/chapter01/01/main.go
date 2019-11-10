package main

import (
	"bytes"
	"fmt"
)

// https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

// Problem: Give an example you can use a composite pattern.

type IComposite interface {
	Kind() string
}

type pig string

func (p pig) Kind() string {
	return fmt.Sprintf("%s: I'm an animal!", p)
}

type dolphine string

func (d dolphine) Kind() string {
	return fmt.Sprintf("%s: I'm a mammal", d)
}

type person struct {
	fname string
	lname string
}

func (p person) Kind() string {
	return fmt.Sprintf("%s: I'm a person", p.fullname())
}

func (p person) fullname() string {
	return fmt.Sprintf("%s %s", p.fname, p.lname)
}

type Composite struct {
	iCompo []IComposite
}

func (c *Composite) Kind() string {
	var buf bytes.Buffer
	for _, compo := range c.iCompo {
		fmt.Fprintf(&buf, "\n-> %s", compo.Kind())
	}
	return fmt.Sprintln("Composite: Hey! I'm composite and I'm composed of:", buf.String())
}

func main() {
	compo := &Composite{
		iCompo: []IComposite{
			pig("Baboy"),
			dolphine("idol"),
			person{"Luffy", "Monkey"},
		},
	}
	fmt.Println(compo.Kind())
}
