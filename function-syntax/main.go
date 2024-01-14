package main

import "fmt"

func main() {
	arg("hello")
	s := arg2("hello")
	fmt.Println("returned value: ", s)
	s1, i := arg3("hello", 10)
	fmt.Println("returned values", s1, i)

}

//Functions Syntax

func arg(s string) {
	fmt.Println("argument", s)
}

func arg2(s string) string {
	return fmt.Sprint("argument:", s)
}

func arg3(s string, i int) (string, int) {
	return fmt.Sprint("some arg:", s), i
}
