package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("this is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}
func loginfo(s fmt.Stringer) {
	log.Println("log from umar ", s.String())
}
func main() {
	b := book{
		title: "West with the night",
	}
	var c count = 42
	loginfo(b)
	loginfo(c)
	// fmt.Println(b)
	// fmt.Println(c)
}
