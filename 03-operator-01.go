package main

import "fmt"

func operatorOne() {
	var (
		x int
		y int
	)

	// x = 7
	// y = 3

	x = 15
	y = 15

	// Arithmetic + - / * %
	fmt.Println("----------- Arithmetic")
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(x % y) // % the remainder of the division

	// Relational == != < > <= >=
	fmt.Println("----------- Relational")
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)

	// Logical && ||
	fmt.Println("----------- Logical")
	fmt.Println(x == y && x != y)
	fmt.Println(x > y && x != y)
	fmt.Println(x == y || x != y)
	fmt.Println(x > y || x != y)
}
