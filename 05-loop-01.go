package main

import "fmt"

func loopOne() {
	/*
		for InitSimpleStatement; Condition; UpdateStatement; {
			code
		}
	*/

	fmt.Println("----------- for")
	// For loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------- while")
	// While loop
	w := 0
	for w < 5 {
		fmt.Println(w)
		w++
	}

	fmt.Println("----------- break")
	// Break loop
	b := 0

	for {
		if b > 6 {
			break
		}
		fmt.Println(b)
		b++
	}

	fmt.Println("----------- continue")
	// Continue loop
	c := 0
	for ; c <= 8; c++ {
		if c%2 == 1 {
			continue
		}
		fmt.Println(c)
	}
}
