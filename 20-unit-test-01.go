package main

import (
	"errors"
	"fmt"
)

func unitTestOne() {
	result, err := divide(15, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func divide(a, b float32) (float32, error) {
	if a == 0 {
		return 0, errors.New("dividende : impossible de diviser par zéro")
	}

	if b == 0 {
		return 0, errors.New("diviseur : impossible de diviser par zéro")
	}

	return a / b, nil
}
