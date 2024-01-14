package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("Hey ", p.first)
}

func main() {
	p1 := person{
		first: "umar",
	}
	p1.speak()
}
