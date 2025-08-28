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

	fmt.Printf("O Tipo de E é %T", e)

}
