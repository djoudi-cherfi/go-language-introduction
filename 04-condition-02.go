package main

import (
	"fmt"
	"math/rand"
	"time"
)

func conditionTwo() {
	// Random seed generator per second
	rand.Seed(time.Now().UnixNano())

	age := 19

	if age > 18 {
		fmt.Println(age, "Je suis majeur !")
	} else if age == 18 {
		fmt.Println(age, "Je viens tout juste d'être majeur !")
	} else {
		fmt.Println(age, "Je suis mineur !")
	}
}
