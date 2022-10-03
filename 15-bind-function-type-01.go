package main

import "fmt"

func bindFunctionType() {
	myContactOne := newContact("John", 30, map[string]string{"Tel": "01.23.45.67.89", "mobile": "06.12.34.56.78"})

	// fmt.Println(myContactOne)
	fmt.Println(myContactOne.displayContact())

	myContactTwo := newContact("Jane", 28, map[string]string{"Tel": "01.67.45.23.89", "mobile": "06.56.34.12.78"})

	// fmt.Println(myContactTwo)
	fmt.Println(myContactTwo.displayContact())

	myContactThree := newContact("Chris", 35, map[string]string{"Tel": "01.02.03.04.05", "mobile": "06.07.08.09.10"})

	// fmt.Println(myContactThree)
	fmt.Println(myContactThree.displayContact())
}
