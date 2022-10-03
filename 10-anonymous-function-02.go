package main

import "fmt"

func anonymousFunctionTwo() {
	fmt.Println("----------- x, y")
	x, y := func() (int, int) {
		fmt.Println("Aucun paramètre. Deux returns")

		return 5, 7
	}()

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("----------- a, b")
	func(a, b int) {
		fmt.Println("A * A + B * B =", a*a+b*b)
	}(x, y)
}
