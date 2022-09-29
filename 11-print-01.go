package main

import "fmt"

func printOne() {
	// Amener à disparaitre dans le futur
	println("Salut à tous !")
	print("Salut à tous !\n")

	// Utiliser fmt
	fmt.Println("Salut à tous !")
	fmt.Print("Salut à tous !\n")

	// Printf
	x, y := 50, 50
	fmt.Printf("Mon nombre (default) est : %v\n", x)
	fmt.Printf("Mon nombre (base 2) est : %b\n", x)
	fmt.Printf("Mon nombre (base 8) est : %o\n", x)
	fmt.Printf("Mon nombre (base 8) est : %O\n", x)
	fmt.Printf("Mon nombre (base 10) est : %v\n", x)
	fmt.Printf("Mon nombre (base 16) est : %x\n", x)
	fmt.Printf("Mon nombre (base 16) est : %X\n", x)
	fmt.Printf("La valeur de X est égale à la valeur de Y => %t\n", x == y)
	fmt.Printf("La valeur de X est égale à la valeur de Y => %T\n", x == y)
	fmt.Printf("L'écriture scientifique de 1258.16559 => %E\n", 1258.16559)
}
