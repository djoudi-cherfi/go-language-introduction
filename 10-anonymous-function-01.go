package main

import "fmt"

func anonymousFunctionOne() {
	fmt.Println("----------- w")
	w := func() {
		fmt.Println("Je suis une fonction anonyme sans return")
	}

	w()

	fmt.Println("----------- without")
	func() {
		fmt.Println("Je suis une fonction anonyme")
	}()

	fmt.Println("----------- z")
	z := func() string {
		fmt.Println("Je suis une fonction anonyme avec un return")

		return "Doe"
	}()

	fmt.Println(z)
}
