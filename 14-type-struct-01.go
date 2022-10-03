package main

import "fmt"

func typeStructOne() {
	type Example struct {
		A string
		B int
		C bool
	}

	myExample := Example{
		A: "John",
		B: 42,
		C: true,
	}

	fooExample := Example{"Jane", 32, true}

	fmt.Println(myExample)
	fmt.Println(fooExample)

	// myContact := newContact("John", 30)
	// fmt.Println(myContact)
}
