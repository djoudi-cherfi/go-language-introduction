package main

import "fmt"

func operatorOne() {
	// Relational == != < <= > >=

	var (
		x int
		y int
	)

	x = 7
	y = 3

	// Arithmetic + - / * %
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(x % y) // % the remainder of the division
}
