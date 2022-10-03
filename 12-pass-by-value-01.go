package main

import (
	"fmt"
)

/*
	Par définition "pass by value" signifie que c'est une copie en mémoire
	de la valeur du paramètre qui est transmis.
*/

func passByValueOne() {
	// A : string, int, bool, float, array
	number := 10

	number = updateA(number)

	fmt.Println(number) // 5

	// B : map, function
	groceryList := map[string]int{
		"gateau": 8,
		"eau":    2,
		"viande": 6,
	}

	updateB(groceryList)

	fmt.Println(groceryList)
}

func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["bonbon"] = 4
	item["tournevis"] = 7
}
