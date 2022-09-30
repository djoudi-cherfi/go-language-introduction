package main

import (
	"fmt"
	"tuto/go/footypes"
)

func moduleOne() {
	var fooVar footypes.Foo

	fooVar.TypeString = "John"
	fooVar.TypeInt = 18
	fooVar.TypeBoolean = true

	fmt.Println(fooVar)
}
