package main

import "fmt"

const a = "Hello, World"

type ID int

var (
	b bool    = true
	c int     = 23
	d string  = "PG"
	e float64 = 2025.2
	f ID      = 1
)

func main() {

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray {

		fmt.Printf("O Valor do Índice é %d e o Valor é %d\n", i, v)
	}
}
