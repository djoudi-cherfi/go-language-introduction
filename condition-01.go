package main

import (
	"fmt"
	"math/rand"
	"time"
)

func conditionOne() {
	/*
		if Statement; Condition {
			code
		} else {
			code
		}
	*/

	// Random seed generator per second
	rand.Seed(time.Now().UnixNano())

	// Random number generator
	fmt.Println(rand.Int())

	// Random number statement with remainder condition return random number even or odd
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, "est un nombre paire !")
	} else {
		fmt.Println(x, "est un nombre impaire !")
	}

	// Remainder variable of a random number that returns a remainder even or odd
	y := rand.Int() % 2

	if y == 0 {
		fmt.Println(y, "est paire!")
	} else {
		fmt.Println(y, "est impaire!")
	}
}
