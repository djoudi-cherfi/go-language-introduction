package main

import "fmt"

var x int

func scopeOne() {
	x = 5
	fmt.Println(x) // 5
	f()            // 10
}

func f() {
	x := 10
	fmt.Println(x)
}
