package main

import "fmt"

func main() {
	y := foo()
	fmt.Println(y())
}

// return a function
func foo() func() int {
	return func() int {
		return 43
	}
}
