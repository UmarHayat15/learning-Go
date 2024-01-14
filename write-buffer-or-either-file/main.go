package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	p := person{
		first: "umar",
	}

	f, err := os.Create("Output.txt")
	if err != nil {
		log.Fatalf("error %s \n", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
