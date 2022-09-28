package main

import "fmt"

func variableOne() {
	var x int // Type after the identifier
	x = 16    // Variable assigned after declaration
	y := 17   // Variable declared and assigned

	fmt.Printf("Variable x = %#v et y = %#v\n", x, y)

	var (
		a int
		b string
		c bool
	)

	a = 100
	b = "I am a string"
	c = true

	fmt.Printf("a = %#v, b = %#v, c = %#v\n", a, b, c)
}
