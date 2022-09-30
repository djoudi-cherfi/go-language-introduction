package main

import "fmt"

func pointerOne() {
	number := 10

	fmt.Println(number)                                 // 10
	fmt.Println("Adresse mémoire de number :", &number) // 0xc0000...

	myPointer := &number // Pointer

	fmt.Println(myPointer) // 0xc0000...
	fmt.Printf("La valeur de l'adresse mémoire %v est %d\n", myPointer, *myPointer)

}
