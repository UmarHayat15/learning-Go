package main

import (
	"bytes"
	"fmt"
)

// create buffer and perform read write operation on it
func main() {
	b := bytes.NewBufferString("Hello")
	fmt.Println(b.String())
	b.WriteString("Umar")
	fmt.Println(b.String())
	b.Reset()
	fmt.Print(b.String())
	b.Write([]byte("umar"))
	fmt.Println(b.String())
}
