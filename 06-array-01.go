package main

import "fmt"

func arrayOne() {
	/*
		var list [n]type
		newList := [...]type{1, 2, 3...}
	*/

	fmt.Println("----------- list")
	var list [2]int

	list[0] = 10
	list[1] = 20

	fmt.Println(list[0])
	fmt.Println(list[1])

	fmt.Println("----------- newList")
	newList := [...]int{40, 50}

	fmt.Println(newList[0])
	fmt.Println(newList[1])

	fmt.Println("----------- bigList")
	bigList := [...]int{60, 80, 200, 800, 500000, 999999}

	for k, v := range bigList {
		fmt.Printf("Position %d est égale à %d.\n", k, v)
	}
}
