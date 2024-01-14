package main

import "fmt"

func main() {
	i := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("sum:", i)
}

// variadic parameters
func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}
