package main

import "fmt"

type person struct {
	first string
}
type agent struct {
	person
	kl bool
}

func (p person) speak() {
	fmt.Println("Hey ", p.first)
}

func (a agent) speak() {
	fmt.Println("Hemlo agent", a.first)
}

type human interface {
	speak()
}

func say(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "umar",
	}
	p2 := agent{
		person{
			first: "hayat",
		},
		false,
	}
	say(p1)
	say(p2)
}
