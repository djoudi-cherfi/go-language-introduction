package main

import "fmt"

type cat struct {
	name  string
	breed string
}

type spider struct {
	name     string
	breed    string
	venomous bool
}

type Animal interface {
	Noise() string
	NumberOfLegs() int
}

func (c *cat) Noise() string {
	return "miaou"
}

func (c *cat) NumberOfLegs() int {
	return 4
}

func (s *spider) Noise() string {
	return "hiss"
}

func (s *spider) NumberOfLegs() int {
	return 8
}

func printAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.Noise(), "est il possède", a.NumberOfLegs(), "pattes !")
}

func interfaceOne() {
	catOne := cat{
		name:  "Kitty",
		breed: "Siamois",
	}

	printAnimalInfo(&catOne)

	spiderOne := spider{
		name:     "Spi",
		breed:    "Veuve noir",
		venomous: true,
	}

	printAnimalInfo(&spiderOne)

}
