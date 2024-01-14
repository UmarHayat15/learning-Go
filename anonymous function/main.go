package main

import "fmt"

func main() {
	func() {
		fmt.Println("there is a function")
	}()

	y := func() {
		fmt.Println("Another function")
	}

	y()
}
