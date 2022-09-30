package main

import "fmt"

func typeStructOne() {
	type Example struct {
		a string
		b int
		c bool
	}

	myExample := Example{
		a: "John",
		b: 42,
		c: true,
	}

	fooExample := Example{"Jane", 32, true}

	fmt.Println(myExample)
	fmt.Println(fooExample)

	// myContact := newContact("John", 30)
	// fmt.Println(myContact)
}
