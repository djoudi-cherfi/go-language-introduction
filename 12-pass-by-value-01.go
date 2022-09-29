package main

import (
	"fmt"
)

/*
	Par définition "pass by value" signifie que c'est une copie en mémoire
	de la valeur du paramètre qui est transmis.
*/

func passByValue() {
	// A : string, int, bool, float, array
	number := 10

	updateA(number)

	fmt.Println(number) // 10
}

func updateA(number int) {
	number = 5
}
