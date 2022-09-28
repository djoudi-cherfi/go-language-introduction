package main

import (
	"errors"
	"fmt"
)

func functionOne() {
	sayHello("John")

	name, err := sayBye("Kate")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)
	}

	result := calcPeriRect(5, 10)
	fmt.Printf("Le périmètre de mon rectangle est : %v.\n", result)
}

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v.\n", name)
}

// func calcPeriRect(length int, width int) {}
func calcPeriRect(length, width int) int {
	return 2 * (length + width)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("erreur: Vous avez oublié de spécifier un nom")
	}

	byeMessage := fmt.Sprintf("%v s'en va ! Bonne soirée.", name)

	return byeMessage, nil
}
