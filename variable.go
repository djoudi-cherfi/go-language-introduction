package main

import "fmt"

func variable() {
	var x int // Type after the identifier
	x = 16    // Variable assigned after declaration
	y := 17   // Variable declared and assigned

	fmt.Printf("Variable x = %#v et y = %#v\n", x, y)
}
